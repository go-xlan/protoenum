package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/stretchr/testify/require"
)

// TestNewEnum tests the creation and basic methods of Enum instance
// Checks Code, Name, Basic, and Meta methods return expected values
//
// 验证 Enum 包装器的创建和基本方法
// 测试 Code、Name、Basic 和 Meta 方法返回预期值
func TestNewEnum(t *testing.T) {
	type StatusType string
	const (
		StatusTypeSuccess StatusType = "success"
	)

	enum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "任务完成")
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())
	t.Log(enum.Meta().Desc())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Basic(), StatusTypeSuccess)
	require.Equal(t, enum.Meta().Desc(), "任务完成")
}

// TestEnum_Proto tests the Proto method returns the source enum
// Checks that the base Protocol Buffer enum is accessible and unchanged
//
// 验证 Proto 方法返回原始枚举
// 测试底层 Protocol Buffer 枚举可访问且未改变
func TestEnum_Proto(t *testing.T) {
	type StatusType string
	const (
		StatusTypeSuccess StatusType = "success"
	)

	enum := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	statusEnum := enum.Proto()

	t.Log(statusEnum.String())
	t.Log(statusEnum.Number())
	require.Equal(t, statusEnum.String(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, statusEnum.Number(), protoenumstatus.StatusEnum_SUCCESS.Number())

	t.Log(statusEnum.Type().Descriptor().Name())
	require.Equal(t, statusEnum.Type().Descriptor().Name(), protoenumstatus.StatusEnum_SUCCESS.Type().Descriptor().Name())
}

// TestEnum_Basic tests the Basic method returns the Go native enum value
// Checks that the basic enum value is accessible and matches the source
//
// 验证 Basic 方法返回 Go 原生枚举值
// 测试 basic 枚举值可访问且与原始值匹配
func TestEnum_Basic(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enumUnknown := protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown)
	enumSuccess := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess)
	enumFailure := protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure)

	t.Log(enumUnknown.Basic())
	t.Log(enumSuccess.Basic())
	t.Log(enumFailure.Basic())

	require.Equal(t, enumUnknown.Basic(), StatusTypeUnknown)
	require.Equal(t, enumSuccess.Basic(), StatusTypeSuccess)
	require.Equal(t, enumFailure.Basic(), StatusTypeFailure)

	// Check Basic returns the exact type
	// 验证 Basic 返回精确的类型
	basicValue := enumSuccess.Basic()
	require.Equal(t, basicValue, StatusTypeSuccess)
	require.Equal(t, string(basicValue), "success")
}

// TestNewEnumWithMeta tests custom metadata type with NewEnumWithMeta
// Checks that custom meta types work with the Enum instance
//
// 验证 NewEnumWithMeta 支持自定义元数据类型
// 测试自定义 meta 类型与 Enum 包装器配合工作
func TestNewEnumWithMeta(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	// MetaI18n represents a custom metadata type with English/Chinese descriptions
	// MetaI18n 代表带有中英文描述的自定义元数据类型
	type MetaI18n struct {
		zh string // Chinese description // 中文描述
		en string // English description // 英文描述
	}

	// Create enums with custom English/Chinese metadata
	// 使用自定义中英文元数据创建枚举
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

	// Check Basic works with custom meta as expected
	// 验证 Basic 在自定义 meta 下仍然工作
	require.Equal(t, StatusTypeSuccess, enumSuccess.Basic())
	require.Equal(t, int32(protoenumstatus.StatusEnum_SUCCESS.Number()), enumSuccess.Code())
}
