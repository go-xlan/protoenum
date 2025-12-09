package main

import (
	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// ResultType represents a Go native enum of result
// ResultType 代表结果的 Go 原生枚举
type ResultType string

const (
	ResultTypeUnknown ResultType = "unknown"
	ResultTypePass    ResultType = "pass"
	ResultTypeMiss    ResultType = "miss"
	ResultTypeSkip    ResultType = "skip"
)

// Build enum collection with description
// 构建带描述的枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown, "其它"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_PASS, ResultTypePass, "通过"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_MISS, ResultTypeMiss, "出错"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_SKIP, ResultTypeSkip, "跳过"),
)

func main() {
	// Lookup using enum code (returns default when not found)
	// 按枚举代码查找（找不到时返回默认值）
	skip := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	zaplog.LOG.Debug("basic", zap.String("msg", string(skip.Basic())))
	zaplog.LOG.Debug("desc", zap.String("msg", skip.Meta().Desc()))

	// Lookup using Go native enum value (type-safe)
	// 按 Go 原生枚举值查找（类型安全查找）
	pass := enums.GetByBasic(ResultTypePass)
	base := protoenumresult.ResultEnum(pass.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Business logic with native enum
	// 使用原生枚举的业务逻辑
	if base == protoenumresult.ResultEnum_PASS {
		zaplog.LOG.Debug("pass")
	}

	// Lookup using enum name (safe with default fallback)
	// 按枚举名称查找（安全且有默认值回退）
	miss := enums.GetByName("MISS")
	zaplog.LOG.Debug("basic", zap.String("msg", string(miss.Basic())))
	zaplog.LOG.Debug("desc", zap.String("msg", miss.Meta().Desc()))

	// List each basic enum value in defined sequence
	// 按定义次序列出各 basic 枚举值
	basics := enums.ListBasics()
	for _, basic := range basics {
		zaplog.LOG.Debug("list", zap.String("basic", string(basic)))
	}
}
