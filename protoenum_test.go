package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/stretchr/testify/require"
)

// TestNewEnum tests the creation and basic methods of Enum instance
// Checks Code, Name, Pure, and Meta methods return expected values
//
// 验证 Enum 包装器的创建和基本方法
// 测试 Code、Name、Pure 和 Meta 方法返回预期值
func TestNewEnum(t *testing.T) {
	type StatusType string
	const (
		StatusTypeSuccess StatusType = "success"
	)

	enum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "任务完成")
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Pure())
	t.Log(enum.Meta().Desc())
	t.Log(enum.Meta().Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Pure(), StatusTypeSuccess)
	require.Equal(t, enum.Meta().Desc(), "任务完成")
	require.Equal(t, enum.Meta().Hans(), "任务完成")
}

// TestEnum_Base tests the Base method returns the source enum
// Checks that the base Protocol Buffer enum is accessible and unchanged
//
// 验证 Base 方法返回原始枚举
// 测试底层 Protocol Buffer 枚举可访问且未改变
func TestEnum_Base(t *testing.T) {
	type StatusType string
	const (
		StatusTypeSuccess StatusType = "success"
	)

	enum := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Pure())

	statusEnum := enum.Base()

	t.Log(statusEnum.String())
	t.Log(statusEnum.Number())
	require.Equal(t, statusEnum.String(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, statusEnum.Number(), protoenumstatus.StatusEnum_SUCCESS.Number())

	t.Log(statusEnum.Type().Descriptor().Name())
	require.Equal(t, statusEnum.Type().Descriptor().Name(), protoenumstatus.StatusEnum_SUCCESS.Type().Descriptor().Name())
}

// TestEnum_Pure tests the Pure method returns the Go native enum value
// Checks that the plain enum value is accessible and matches the source
//
// 验证 Pure 方法返回 Go 原生枚举值
// 测试朴素枚举值可访问且与原始值匹配
func TestEnum_Pure(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enumUnknown := protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown)
	enumSuccess := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess)
	enumFailure := protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure)

	t.Log(enumUnknown.Pure())
	t.Log(enumSuccess.Pure())
	t.Log(enumFailure.Pure())

	require.Equal(t, enumUnknown.Pure(), StatusTypeUnknown)
	require.Equal(t, enumSuccess.Pure(), StatusTypeSuccess)
	require.Equal(t, enumFailure.Pure(), StatusTypeFailure)

	// Check Pure returns the exact type
	// 验证 Pure 返回精确的类型
	pureValue := enumSuccess.Pure()
	require.Equal(t, pureValue, StatusTypeSuccess)
	require.Equal(t, string(pureValue), "success")
}

// TestNewEnumWithMeta tests custom metadata type with NewEnumWithMeta
// Checks that user-defined meta types work with the Enum instance
//
// 验证 NewEnumWithMeta 支持自定义元数据类型
// 测试用户自定义的 meta 类型与 Enum 包装器配合工作
func TestNewEnumWithMeta(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	// MetaI18n represents a custom metadata type with bilingual descriptions
	// MetaI18n 代表带有双语描述的自定义元数据类型
	type MetaI18n struct {
		zh string // Chinese description // 中文描述
		en string // English description // 英文描述
	}

	// Create enums with custom bilingual metadata
	// 使用自定义双语元数据创建枚举
	enumUnknown := protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, &MetaI18n{zh: "未知", en: "Unknown"})
	enumSuccess := protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, &MetaI18n{zh: "成功", en: "Success"})
	enumFailure := protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure, &MetaI18n{zh: "失败", en: "Failure"})

	t.Log(enumUnknown.Meta().zh, enumUnknown.Meta().en)
	t.Log(enumSuccess.Meta().zh, enumSuccess.Meta().en)
	t.Log(enumFailure.Meta().zh, enumFailure.Meta().en)

	// Check custom meta fields
	// 验证自定义 meta 字段
	require.Equal(t, "未知", enumUnknown.Meta().zh)
	require.Equal(t, "Unknown", enumUnknown.Meta().en)
	require.Equal(t, "成功", enumSuccess.Meta().zh)
	require.Equal(t, "Success", enumSuccess.Meta().en)
	require.Equal(t, "失败", enumFailure.Meta().zh)
	require.Equal(t, "Failure", enumFailure.Meta().en)

	// Check Pure still works with custom meta
	// 验证 Pure 在自定义 meta 下仍然工作
	require.Equal(t, StatusTypeSuccess, enumSuccess.Pure())
	require.Equal(t, int32(protoenumstatus.StatusEnum_SUCCESS.Number()), enumSuccess.Code())
}
