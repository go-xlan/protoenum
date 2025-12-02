package main

import (
	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// StatusType represents a Go native enum of status
// StatusType 代表状态的 Go 原生枚举
type StatusType string

const (
	StatusTypeUnknown StatusType = "unknown"
	StatusTypeSuccess StatusType = "success"
	StatusTypeFailure StatusType = "failure"
)

// MetaI18n represents a custom metadata type with dual-language descriptions
// MetaI18n 代表带有双语描述的自定义元数据类型
type MetaI18n struct {
	zhCN string // Chinese description // 中文描述
	enUS string // English description // 英文描述
}

func (c *MetaI18n) Chinese() string { return c.zhCN }
func (c *MetaI18n) English() string { return c.enUS }

// Build enum collection with custom dual-language metadata
// 构建带有自定义双语元数据的枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, &MetaI18n{zhCN: "未知", enUS: "Unknown"}),
	protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, &MetaI18n{zhCN: "成功", enUS: "Success"}),
	protoenum.NewEnumWithMeta(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure, &MetaI18n{zhCN: "失败", enUS: "Failure"}),
)

func main() {
	// Lookup using Go native enum value (type-safe)
	// 按 Go 原生枚举值查找（类型安全查找）
	success := enums.GetByPure(StatusTypeSuccess)
	zaplog.LOG.Debug("pure", zap.String("msg", string(success.Pure())))
	zaplog.LOG.Debug("zh-CN", zap.String("msg", success.Meta().Chinese()))
	zaplog.LOG.Debug("en-US", zap.String("msg", success.Meta().English()))

	// Lookup using enum code (returns default when not found)
	// 按枚举代码查找（找不到时返回默认值）
	failure := enums.GetByCode(int32(protoenumstatus.StatusEnum_FAILURE))
	zaplog.LOG.Debug("pure", zap.String("msg", string(failure.Pure())))
	zaplog.LOG.Debug("zh-CN", zap.String("msg", failure.Meta().Chinese()))
	zaplog.LOG.Debug("en-US", zap.String("msg", failure.Meta().English()))

	// Business logic with native enum
	// 使用原生枚举的业务逻辑
	if success.Pure() == StatusTypeSuccess {
		zaplog.LOG.Debug("done")
	}
}
