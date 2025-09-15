package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/stretchr/testify/require"
)

// enums is a test fixture containing each status enum value
// Provides a complete Enums collection when testing each lookup method
//
// enums 是包含所有状态枚举值的测试固定装置
// 提供完整的 Enums 集合用于测试所有查找方法
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)

// TestEnums_GetByEnum verifies lookup using Protocol Buffer enum value
// Tests that GetByEnum returns the correct Enum with matching properties
//
// 验证通过 Protocol Buffer 枚举值检索
// 测试 GetByEnum 返回具有匹配属性的正确 Enum
func TestEnums_GetByEnum(t *testing.T) {
	enum := enums.GetByEnum(protoenumstatus.StatusEnum_SUCCESS)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Hans(), "成功")
}

// TestEnums_GetByCode verifies lookup using numeric code
// Tests that GetByCode returns the correct Enum using int32 code
//
// 验证通过数字代码检索
// 测试 GetByCode 使用 int32 代码返回正确的 Enum
func TestEnums_GetByCode(t *testing.T) {
	enum := enums.GetByCode(int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_FAILURE.String())
	require.Equal(t, enum.Hans(), "失败")
}

// TestEnums_GetByName verifies lookup using string name
// Tests that GetByName returns the correct Enum using enum name
//
// 验证通过字符串名称检索
// 测试 GetByName 使用枚举名称返回正确的 Enum
func TestEnums_GetByName(t *testing.T) {
	enum := enums.GetByName(protoenumstatus.StatusEnum_UNKNOWN.String())
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_UNKNOWN.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_UNKNOWN.String())
	require.Equal(t, enum.Hans(), "未知")
}

// TestEnums_GetByDesc verifies lookup using description
// Tests that GetByDesc returns the correct Enum using custom description
//
// 验证通过描述检索
// 测试 GetByDesc 使用自定义描述返回正确的 Enum
func TestEnums_GetByDesc(t *testing.T) {
	enum := enums.GetByDesc("成功")
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Hans(), "成功")
}

// TestEnums_GetByHans verifies lookup using Chinese description
// Tests that GetByHans returns the correct Enum using Chinese text
//
// 验证通过中文描述检索
// 测试 GetByHans 使用中文文本返回正确的 Enum
func TestEnums_GetByHans(t *testing.T) {
	enum := enums.GetByHans("成功")
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Hans(), "成功")
}

// TestEnums_DefaultValue verifies default value functionality
// Tests that enums return default value when lookup fails
//
// 验证默认值功能
// 测试查找失败时枚举返回默认值
func TestEnums_DefaultValue(t *testing.T) {
	// Test that first item becomes default
	defaultEnum := enums.GetDefault()
	require.NotNil(t, defaultEnum)
	require.Equal(t, "未知", defaultEnum.Desc())
	require.Equal(t, int32(protoenumstatus.StatusEnum_UNKNOWN), defaultEnum.Code())

	// Test lookup returns default when not found
	notFound := enums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, defaultEnum, notFound)

	// Test GetByName returns default when not found
	notFoundByName := enums.GetByName("NOT_EXISTS")
	require.NotNil(t, notFoundByName)
	require.Equal(t, defaultEnum, notFoundByName)

	// Test GetByDesc returns default when not found
	notFoundByDesc := enums.GetByDesc("不存在的描述")
	require.NotNil(t, notFoundByDesc)
	require.Equal(t, defaultEnum, notFoundByDesc)
}

// TestEnums_SetDefault verifies dynamic default value setting
// Tests that SetDefault allows changing the default value after creation
//
// 验证动态设置默认值
// 测试 SetDefault 允许在创建后更改默认值
func TestEnums_SetDefault(t *testing.T) {
	// Create a new enum collection
	customEnums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
	)

	// Change default to SUCCESS
	successEnum := customEnums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	customEnums.SetDefault(successEnum)

	// Verify new default
	newDefault := customEnums.GetDefault()
	require.NotNil(t, newDefault)
	require.Equal(t, "成功", newDefault.Desc())

	// Test lookup returns new default when not found
	notFound := customEnums.GetByCode(999)
	require.NotNil(t, notFound)
	require.Equal(t, "成功", notFound.Desc())
}

// TestEnums_ChainMethods verifies chain-style configuration methods
// Tests WithDefaultEnum, WithDefaultCode, and WithDefaultName for fluent API
// Also verifies panic behavior when invalid code or name is specified
//
// 验证链式配置方法
// 测试 WithDefaultEnum、WithDefaultCode 和 WithDefaultName 的流式 API
// 同时验证指定无效代码或名称时的 panic 行为
func TestEnums_ChainMethods(t *testing.T) {
	// Test WithDefaultEnum chain method
	t.Run("with-default-enum", func(t *testing.T) {
		customEnums := protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
		).WithDefaultEnum(protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"))

		require.NotNil(t, customEnums.GetDefault())
		require.Equal(t, "成功", customEnums.GetDefault().Desc())
	})

	// Test WithDefaultCode chain method
	t.Run("with-default-code", func(t *testing.T) {
		customEnums := protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
		).WithDefaultCode(int32(protoenumstatus.StatusEnum_FAILURE))

		require.NotNil(t, customEnums.GetDefault())
		require.Equal(t, "失败", customEnums.GetDefault().Desc())
	})

	// Test WithDefaultName chain method
	t.Run("with-default-name", func(t *testing.T) {
		customEnums := protoenum.NewEnums(
			protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
			protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
		).WithDefaultName("SUCCESS")

		require.NotNil(t, customEnums.GetDefault())
		require.Equal(t, "成功", customEnums.GetDefault().Desc())
	})

	// Test chain with non-existent code (should panic)
	t.Run("with-invalid-code-panics", func(t *testing.T) {
		require.Panics(t, func() {
			protoenum.NewEnums(
				protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
				protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
				protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
			).WithDefaultCode(999)
		})
	})

	// Test chain with non-existent name (should panic)
	t.Run("with-invalid-name-panics", func(t *testing.T) {
		require.Panics(t, func() {
			protoenum.NewEnums(
				protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
				protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
				protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
			).WithDefaultName("NOT_EXISTS")
		})
	})
}
