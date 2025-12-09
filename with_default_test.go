package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/stretchr/testify/require"
)

// TestEnums_DefaultValue verifies default value features
// Tests that enums return default value when lookup fails
//
// 验证默认值功能
// 测试查找失败时枚举返回默认值
func TestEnums_DefaultValue(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	// Test that first item becomes default
	defaultEnum := enums.GetDefault()
	require.NotNil(t, defaultEnum)
	require.Equal(t, StatusTypeUnknown, defaultEnum.Basic())
	require.Equal(t, int32(protoenumstatus.StatusEnum_UNKNOWN), defaultEnum.Code())

	// Test lookup returns default when not found
	notFound := enums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, defaultEnum, notFound)

	// Test GetByName returns default when not found
	notFoundByName := enums.GetByName("NOT_EXISTS")
	require.NotNil(t, notFoundByName)
	require.Equal(t, defaultEnum, notFoundByName)

	// Test GetByBasic returns default when not found
	notFoundByBasic := enums.GetByBasic(StatusType("not_exists"))
	require.NotNil(t, notFoundByBasic)
	require.Equal(t, defaultEnum, notFoundByBasic)
}

// TestEnums_GetDefaultProto verifies GetDefaultProto returns the protoEnum value
// Tests that GetDefaultProto returns the correct Protocol Buffer enum type
//
// 验证 GetDefaultProto 返回 protoEnum 值
// 测试 GetDefaultProto 返回正确的 Protocol Buffer 枚举类型
func TestEnums_GetDefaultProto(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	// Test GetDefaultProto returns the protoEnum value
	defaultEnum := enums.GetDefaultProto()
	t.Log(defaultEnum)
	require.Equal(t, protoenumstatus.StatusEnum_UNKNOWN, defaultEnum)

	// Test panic when no default is set
	enums.UnsetDefault()
	require.Panics(t, func() {
		enums.GetDefaultProto()
	})
}

// TestEnums_GetDefaultBasic verifies GetDefaultBasic returns the basicEnum value
// Tests that GetDefaultBasic returns the correct Go native enum type
//
// 验证 GetDefaultBasic 返回 basicEnum 值
// 测试 GetDefaultBasic 返回正确的 Go 原生枚举类型
func TestEnums_GetDefaultBasic(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	// Test GetDefaultBasic returns the basicEnum value
	defaultBasic := enums.GetDefaultBasic()
	t.Log(defaultBasic)
	require.Equal(t, StatusTypeUnknown, defaultBasic)

	// Test panic when no default is set
	enums.UnsetDefault()
	require.Panics(t, func() {
		enums.GetDefaultBasic()
	})
}

// TestEnums_SetDefault verifies default value setting once unset
// Tests that SetDefault works once the existing default is unset
//
// 验证丢弃后设置默认值
// 测试 SetDefault 在丢弃现有默认值后可以工作
func TestEnums_SetDefault(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	// Create a new enum collection with auto default
	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	// Must unset default first before setting new one
	enums.UnsetDefault()

	// Now set default to SUCCESS
	successEnum := enums.MustGetByBasic(StatusTypeSuccess)
	enums.SetDefault(successEnum)

	// Check new default
	newDefault := enums.GetDefault()
	require.NotNil(t, newDefault)
	require.Equal(t, StatusTypeSuccess, newDefault.Basic())

	// Test lookup returns new default when not found
	notFound := enums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, StatusTypeSuccess, notFound.Basic())
}

// TestEnums_SetDefaultPanicsOnDuplicate verifies SetDefault panics when default exists
// Tests that SetDefault panics when default value is present
//
// 验证 SetDefault 在已有默认值时 panic
// 测试当默认值已存在时 SetDefault 会 panic
func TestEnums_SetDefaultPanicsOnDuplicate(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	)

	// SetDefault should panic because default is present (first item)
	require.Panics(t, func() {
		successEnum := enums.MustGetByBasic(StatusTypeSuccess)
		enums.SetDefault(successEnum)
	})
}

// TestEnums_SetDefaultNilPanics verifies that SetDefault panics with nil input
// Tests that passing nil to SetDefault causes panic
//
// 验证 SetDefault 在 nil 参数时 panic
// 测试传递 nil 给 SetDefault 会导致 panic
func TestEnums_SetDefaultNilPanics(t *testing.T) {
	type StatusType string

	// Create blank collection without default
	enums := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]()

	// SetDefault with nil should panic (must.Full check)
	require.Panics(t, func() {
		enums.SetDefault(nil)
	})
}

// TestEnums_UnsetDefault verifies unsetting the default value
// Tests that UnsetDefault removes the default value and lookups panic
//
// 验证取消设置默认值
// 测试 UnsetDefault 移除默认值后查找会 panic
func TestEnums_UnsetDefault(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	// Create a new enum collection with default
	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	// Check default exists
	require.NotNil(t, enums.GetDefault())
	require.Equal(t, StatusTypeUnknown, enums.GetDefault().Basic())

	// Unset the default
	enums.UnsetDefault()

	// Check GetDefault panics once unset
	require.Panics(t, func() {
		enums.GetDefault()
	})

	// Test GetByCode also panics when not found (because no default)
	require.Panics(t, func() {
		enums.GetByCode(999)
	})
}

// TestEnums_WithUnsetDefault verifies chain-style default unset
// Tests that WithUnsetDefault removes default value and returns the instance
//
// 验证链式取消设置默认值
// 测试 WithUnsetDefault 移除默认值并返回实例
func TestEnums_WithUnsetDefault(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	// Create enum collection and unset default in chain
	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	).WithUnsetDefault()

	// Check GetDefault panics once chain unset
	require.Panics(t, func() {
		enums.GetDefault()
	})

	// Test GetByCode also panics when not found (because no default)
	require.Panics(t, func() {
		enums.GetByCode(999)
	})
}

// TestEnums_ChainMethods verifies chain-style configuration methods
// Tests WithDefault, WithDefaultCode, and WithDefaultName with fluent API
// Checks panic actions when invalid code and name is specified
//
// 验证链式配置方法
// 测试 WithDefault、WithDefaultCode 和 WithDefaultName 的流式 API
// 同时验证指定无效代码或名称时的 panic 行为
func TestEnums_ChainMethods(t *testing.T) {
	// Test WithDefault chain method - add enum and set as default in one chain
	t.Run("with-default-enum", func(t *testing.T) {
		type StatusType string
		const (
			StatusTypeSuccess StatusType = "success"
		)

		// Create blank collection, then add enum and set as default using chain method
		enums := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]().
			WithDefault(protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess))

		require.NotNil(t, enums.GetDefault())
		require.Equal(t, StatusTypeSuccess, enums.GetDefault().Basic())
	})

	// Test that WithDefaultXxx panics when default is present
	t.Run("with-default-panics-on-existing", func(t *testing.T) {
		type StatusType string
		const (
			StatusTypeUnknown StatusType = "unknown"
			StatusTypeSuccess StatusType = "success"
		)

		require.Panics(t, func() {
			// NewEnums sets first item as default, then WithDefaultCode should panic
			protoenum.NewEnums(
				protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
				protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
			).WithDefaultCode(int32(protoenumstatus.StatusEnum_SUCCESS))
		})
	})

	// Test unset then set pattern
	t.Run("unset-then-set-default", func(t *testing.T) {
		type StatusType string
		const (
			StatusTypeUnknown StatusType = "unknown"
			StatusTypeSuccess StatusType = "success"
			StatusTypeFailure StatusType = "failure"
		)

		enums := protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
			protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
		)
		// Must unset first, then set new default
		enums.UnsetDefault()
		successEnum := enums.MustGetByBasic(StatusTypeSuccess)
		enums.SetDefault(successEnum)

		require.NotNil(t, enums.GetDefault())
		require.Equal(t, StatusTypeSuccess, enums.GetDefault().Basic())
	})

	// Test chain with non-existent code (should panic)
	t.Run("with-invalid-code-panics", func(t *testing.T) {
		type StatusType string

		require.Panics(t, func() {
			protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]().WithDefaultCode(999)
		})
	})

	// Test chain with non-existent name (should panic)
	t.Run("with-invalid-name-panics", func(t *testing.T) {
		type StatusType string

		require.Panics(t, func() {
			protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]().WithDefaultName("NOT_EXISTS")
		})
	})
}
