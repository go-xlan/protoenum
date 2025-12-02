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

// Build status enum collection
// 构建状态枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
)

func main() {
	// Get Go native enum from protobuf enum (returns default when not found)
	// 从 protobuf 枚举获取 Go 原生枚举（找不到时返回默认值）
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("pure", zap.String("msg", string(item.Pure())))

	// Convert between protoenum and native enum (safe with default fallback)
	// 在 protoenum 和原生枚举之间转换（安全且有默认值回退）
	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Use in business logic
	// 在业务逻辑中使用
	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}

	// Get default plain enum value (first item becomes default)
	// 获取默认朴素枚举值（第一个元素成为默认值）
	defaultPure := enums.GetDefaultPure()
	zaplog.LOG.Debug("default", zap.String("msg", string(defaultPure)))
}
