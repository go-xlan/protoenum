package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/stretchr/testify/require"
)

// TestEnums_GetByEnum tests lookup with Protocol Buffer enum value
// Checks that GetByEnum returns the correct Enum with matching properties
//
// 验证通过 Protocol Buffer 枚举值检索
// 测试 GetByEnum 返回具有匹配属性的正确 Enum
func TestEnums_GetByEnum(t *testing.T) {
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

	enum := enums.GetByEnum(protoenumstatus.StatusEnum_SUCCESS)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Pure())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Pure(), StatusTypeSuccess)
}

// TestEnums_GetByCode tests lookup using numeric code
// Checks that GetByCode returns the correct Enum using int32 code
//
// 验证通过数字代码检索
// 测试 GetByCode 使用 int32 代码返回正确的 Enum
func TestEnums_GetByCode(t *testing.T) {
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

	enum := enums.GetByCode(int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Pure())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_FAILURE.String())
	require.Equal(t, enum.Pure(), StatusTypeFailure)
}

// TestEnums_GetByName tests lookup using string name
// Checks that GetByName returns the correct Enum using enum name
//
// 验证通过字符串名称检索
// 测试 GetByName 使用枚举名称返回正确的 Enum
func TestEnums_GetByName(t *testing.T) {
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

	enum := enums.GetByName(protoenumstatus.StatusEnum_UNKNOWN.String())
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Pure())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_UNKNOWN.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_UNKNOWN.String())
	require.Equal(t, enum.Pure(), StatusTypeUnknown)
}

// TestEnums_GetByPure verifies lookup using Go native enum value
// Tests that GetByPure returns the correct Enum using plain enum type
//
// 验证通过 Go 原生枚举值检索
// 测试 GetByPure 使用朴素枚举类型返回正确的 Enum
func TestEnums_GetByPure(t *testing.T) {
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

	// Lookup by Go native enum value
	enum := enums.GetByPure(StatusTypeSuccess)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Pure())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Pure(), StatusTypeSuccess)

	// Verify MustGetByPure works as well
	enumFailure := enums.MustGetByPure(StatusTypeFailure)
	require.Equal(t, enumFailure.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enumFailure.Pure(), StatusTypeFailure)
}

// TestEnums_DefaultValue verifies default value functionality
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
	require.Equal(t, StatusTypeUnknown, defaultEnum.Pure())
	require.Equal(t, int32(protoenumstatus.StatusEnum_UNKNOWN), defaultEnum.Code())

	// Test lookup returns default when not found
	notFound := enums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, defaultEnum, notFound)

	// Test GetByName returns default when not found
	notFoundByName := enums.GetByName("NOT_EXISTS")
	require.NotNil(t, notFoundByName)
	require.Equal(t, defaultEnum, notFoundByName)

	// Test GetByPure returns default when not found
	notFoundByPure := enums.GetByPure(StatusType("not_exists"))
	require.NotNil(t, notFoundByPure)
	require.Equal(t, defaultEnum, notFoundByPure)
}

// TestEnums_SetDefault verifies default value setting after unset
// Tests that SetDefault works after unsetting the existing default
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
	successEnum := enums.MustGetByPure(StatusTypeSuccess)
	enums.SetDefault(successEnum)

	// Check new default
	newDefault := enums.GetDefault()
	require.NotNil(t, newDefault)
	require.Equal(t, StatusTypeSuccess, newDefault.Pure())

	// Test lookup returns new default when not found
	notFound := enums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, StatusTypeSuccess, notFound.Pure())
}

// TestEnums_SetDefaultPanicsOnDuplicate verifies SetDefault panics when default exists
// Tests that SetDefault panics if default value already exists
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

	// SetDefault should panic because default already exists (first item)
	require.Panics(t, func() {
		successEnum := enums.MustGetByPure(StatusTypeSuccess)
		enums.SetDefault(successEnum)
	})
}

// TestEnums_SetDefaultNilPanics verifies that SetDefault panics with nil parameter
// Tests that passing nil to SetDefault causes panic
//
// 验证 SetDefault 在 nil 参数时 panic
// 测试传递 nil 给 SetDefault 会导致 panic
func TestEnums_SetDefaultNilPanics(t *testing.T) {
	type StatusType string

	// Create empty collection without default
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
	require.Equal(t, StatusTypeUnknown, enums.GetDefault().Pure())

	// Unset the default
	enums.UnsetDefault()

	// Check GetDefault panics after unset
	require.Panics(t, func() {
		enums.GetDefault()
	})

	// Test GetByCode also panics when not found (because no default)
	require.Panics(t, func() {
		enums.GetByCode(999)
	})
}

// TestEnums_WithUnsetDefault verifies chain-style default removal
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

	// Check GetDefault panics after chain unset
	require.Panics(t, func() {
		enums.GetDefault()
	})

	// Test GetByCode also panics when not found (because no default)
	require.Panics(t, func() {
		enums.GetByCode(999)
	})
}

// TestEnums_ChainMethods verifies chain-style configuration methods
// Tests WithDefaultEnum, WithDefaultCode, and WithDefaultName with fluent API
// Also verifies panic behavior when invalid code or name is specified
//
// 验证链式配置方法
// 测试 WithDefaultEnum、WithDefaultCode 和 WithDefaultName 的流式 API
// 同时验证指定无效代码或名称时的 panic 行为
func TestEnums_ChainMethods(t *testing.T) {
	// Test WithDefaultEnum chain method - add enum and set as default in one chain
	t.Run("with-default-enum", func(t *testing.T) {
		type StatusType string
		const (
			StatusTypeSuccess StatusType = "success"
		)

		// Create empty collection, then add enum and set as default using chain method
		enums := protoenum.NewEnums[protoenumstatus.StatusEnum, StatusType, *protoenum.MetaNone]().
			WithDefaultEnum(protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess))

		require.NotNil(t, enums.GetDefault())
		require.Equal(t, StatusTypeSuccess, enums.GetDefault().Pure())
	})

	// Test that WithDefaultXxx panics when default already exists
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
		successEnum := enums.MustGetByPure(StatusTypeSuccess)
		enums.SetDefault(successEnum)

		require.NotNil(t, enums.GetDefault())
		require.Equal(t, StatusTypeSuccess, enums.GetDefault().Pure())
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

func TestEnums_ListEnums(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	var enums = protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	protoEnums := enums.ListEnums()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 4)
	require.Equal(t, protoenumresult.ResultEnum_UNKNOWN, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[2])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[3])
}

func TestEnums_ListPures(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	var enums = protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	plainEnums := enums.ListPures()
	t.Log(plainEnums)
	require.Len(t, plainEnums, 4)
	require.Equal(t, ResultTypeUnknown, plainEnums[0])
	require.Equal(t, ResultTypePass, plainEnums[1])
	require.Equal(t, ResultTypeMiss, plainEnums[2])
	require.Equal(t, ResultTypeSkip, plainEnums[3])
}
