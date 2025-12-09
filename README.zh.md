[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/protoenum)](https://pkg.go.dev/github.com/go-xlan/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/protoenum/main.svg)](https://coveralls.io/github/go-xlan/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.23--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/protoenum.svg)](https://github.com/go-xlan/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/protoenum)](https://goreportcard.com/report/github.com/go-xlan/protoenum)

# protoenum

`protoenum` 是一个 Go 语言包，提供管理 Protobuf 枚举元数据的工具。它通过 `Basic()` 方法桥接 Protobuf 枚举和 Go 原生枚举（`type StatusType string`），并提供枚举集合支持简单的代码、名称和 Basic 值查找。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **智能枚举管理**：将 Protobuf 枚举与 Go 原生枚举和自定义元数据包装
🔗 **Go 原生枚举桥接**：通过 `Basic()` 方法无缝转换到 Go 原生枚举类型
⚡ **多方式查找**：支持代码、名称和 Basic 值快速查找
🔄 **类型安全操作**：三泛型保持 protobuf、Go 原生枚举和元数据的类型安全
🛡️ **严格设计**：单一使用模式防止误用，强制要求默认值
🌍 **生产级别**：经过实战检验的企业级枚举处理方案

## 安装

```bash
go get github.com/go-xlan/protoenum
```

## 快速开始

### 定义 Proto 枚举

项目包含示例 proto 文件：
- `protoenumstatus.proto` - 基础状态枚举
- `protoenumresult.proto` - 测试结果枚举

### 基础集合使用

```go
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
	zaplog.LOG.Debug("basic", zap.String("msg", string(item.Basic())))

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

	// Get default basic enum value (first item becomes default)
	// 获取默认 basic 枚举值（第一个元素成为默认值）
	defaultBasic := enums.GetDefaultBasic()
	zaplog.LOG.Debug("default", zap.String("msg", string(defaultBasic)))
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 高级查找方法

```go
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
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)


## API 参考

### 单个枚举操作

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnum(protoEnum, basicEnum)` | 创建枚举实例（无元数据） | `*Enum[P, B, *MetaNone]` |
| `NewEnumWithDesc(protoEnum, basicEnum, desc)` | 创建枚举实例（带描述） | `*Enum[P, B, *MetaDesc]` |
| `NewEnumWithMeta(protoEnum, basicEnum, meta)` | 创建枚举实例（带自定义元数据） | `*Enum[P, B, M]` |
| `enum.Proto()` | 获取底层 protobuf 枚举 | `P` |
| `enum.Code()` | 获取数值代码 | `int32` |
| `enum.Name()` | 获取枚举名称 | `string` |
| `enum.Basic()` | 获取 Go 原生枚举值 | `B` |
| `enum.Meta()` | 获取自定义元数据 | `M` |

### 创建集合

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnums(items...)` | 创建集合并严格验证（第一项成为默认值） | `*Enums[P, B, M]` |

### 存在性检查 (Lookup)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.LookupByProto(proto)` | 按 protobuf 枚举查找，检查是否存在 | `(*Enum[P, B, M], bool)` |
| `enums.LookupByCode(code)` | 按代码查找，检查是否存在 | `(*Enum[P, B, M], bool)` |
| `enums.LookupByName(name)` | 按名称查找，检查是否存在 | `(*Enum[P, B, M], bool)` |
| `enums.LookupByBasic(basic)` | 按 Go 原生枚举查找，检查是否存在 | `(*Enum[P, B, M], bool)` |

### 安全访问 (Get)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.GetByProto(proto)` | 按 protobuf 枚举获取（找不到返回默认值，无默认值则 panic） | `*Enum[P, B, M]` |
| `enums.GetByCode(code)` | 按代码获取（找不到返回默认值，无默认值则 panic） | `*Enum[P, B, M]` |
| `enums.GetByName(name)` | 按名称获取（找不到返回默认值，无默认值则 panic） | `*Enum[P, B, M]` |
| `enums.GetByBasic(basic)` | 按 Go 原生枚举获取（找不到返回默认值，无默认值则 panic） | `*Enum[P, B, M]` |

### 严格访问 (MustGet)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.MustGetByProto(proto)` | 严格按 protobuf 枚举获取（找不到则 panic） | `*Enum[P, B, M]` |
| `enums.MustGetByCode(code)` | 严格按代码获取（找不到则 panic） | `*Enum[P, B, M]` |
| `enums.MustGetByName(name)` | 严格按名称获取（找不到则 panic） | `*Enum[P, B, M]` |
| `enums.MustGetByBasic(basic)` | 严格按 Go 原生枚举获取（找不到则 panic） | `*Enum[P, B, M]` |

### 枚举列表 (List)

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.ListProtos()` | 返回各 protoEnum 值的切片 | `[]P` |
| `enums.ListBasics()` | 返回各 basicEnum 值的切片 | `[]B` |
| `enums.ListValidProtos()` | 返回排除默认值的 protoEnum 切片 | `[]P` |
| `enums.ListValidBasics()` | 返回排除默认值的 basicEnum 切片 | `[]B` |

### 默认值管理

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `enums.GetDefault()` | 获取当前默认值（未设置则 panic） | `*Enum[P, B, M]` |
| `enums.GetDefaultProto()` | 获取默认 protoEnum 值（未设置则 panic） | `P` |
| `enums.GetDefaultBasic()` | 获取默认 basicEnum 值（未设置则 panic） | `B` |
| `enums.SetDefault(enum)` | 设置默认值（要求当前无默认值） | `void` |
| `enums.UnsetDefault()` | 移除默认值（要求当前有默认值） | `void` |
| `enums.WithDefault(enum)` | 链式：通过枚举实例设置默认值 | `*Enums[P, B, M]` |
| `enums.WithDefaultCode(code)` | 链式：通过代码设置默认值（找不到则 panic） | `*Enums[P, B, M]` |
| `enums.WithDefaultName(name)` | 链式：通过名称设置默认值（找不到则 panic） | `*Enums[P, B, M]` |
| `enums.WithUnsetDefault()` | 链式：移除默认值 | `*Enums[P, B, M]` |

## 使用示例

### 单个枚举操作

**创建增强枚举包装器：**
```go
type StatusType string
const StatusTypeSuccess StatusType = "success"

statusEnum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "操作成功")
fmt.Printf("代码: %d, 名称: %s, Basic: %s, 描述: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Basic(), statusEnum.Meta().Desc())
```

**访问底层 protobuf 枚举：**
```go
originalEnum := statusEnum.Proto()
if originalEnum == protoenumstatus.StatusEnum_SUCCESS {
    fmt.Println("检测到成功状态")
}
```

### 集合操作

**构建枚举集合：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
    StatusTypeFailure StatusType = "failure"
)

statusEnums := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知状态"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure, "失败"),
)
```

**多种查找方式：**
```go
// 按数字代码查找 - 始终返回有效枚举（找不到返回默认值）
enum := statusEnums.GetByCode(1)
fmt.Printf("找到: %s\n", enum.Meta().Desc())

// 按枚举名称查找 - 保证非 nil
enum = statusEnums.GetByName("SUCCESS")
fmt.Printf("状态: %s\n", enum.Meta().Desc())

// 按 Go 原生枚举值查找 - 类型安全查找
enum = statusEnums.GetByBasic(StatusTypeSuccess)
fmt.Printf("Basic: %s\n", enum.Basic())

// 严格按 Go 原生枚举值查找（找不到则 panic）
enum = statusEnums.MustGetByCode(1)
fmt.Printf("严格: %s\n", enum.Meta().Desc())
```

**列出枚举值:**
```go
// 获取各已注册 proto 枚举的切片
protoEnums := statusEnums.ListProtos()
// > [UNKNOWN, SUCCESS, FAILURE]

// 获取各已注册 basic Go 原生枚举的切片
basicEnums := statusEnums.ListBasics()
// > ["unknown", "success", "failure"]

// 获取有效值（排除默认值）
validProtos := statusEnums.ListValidProtos()
// > [SUCCESS, FAILURE]（UNKNOWN 是默认值，被排除）

validBasics := statusEnums.ListValidBasics()
// > ["success", "failure"]
```

### 高级用法

**通过 Basic() 桥接 Go 原生枚举：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// 桥接 protobuf 枚举到 Go 原生枚举
enum := enums.GetByCode(1)
basicValue := enum.Basic()  // 返回 StatusType("success")

// 在业务逻辑中使用 Go 原生枚举
switch basicValue {
case StatusTypeSuccess:
    fmt.Println("操作成功")
case StatusTypeUnknown:
    fmt.Println("未知状态")
}

// 通过 Go 原生枚举值查找
found := enums.GetByBasic(StatusTypeSuccess)
fmt.Printf("代码: %d, 名称: %s\n", found.Code(), found.Name())
```

**类型转换模式：**
```go
// 从枚举包装器转换为原生 protobuf 枚举
// 始终返回有效枚举（带默认值回退）
statusEnum := enums.GetByName("SUCCESS")
native := protoenumstatus.StatusEnum(statusEnum.Code())
// 在 protobuf 操作中安全使用原生枚举
```

**严格验证模式：**
```go
// 使用 MustGetByXxx 进行严格验证（找不到则 panic）
result := enums.MustGetByCode(1)  // 找不到会 panic
fmt.Printf("找到: %s\n", result.Name())

// GetByXxx 对未知值返回默认值
result = enums.GetByCode(999)  // 返回默认值（UNKNOWN）
fmt.Printf("回退: %s\n", result.Name())
```

### 默认值和链式配置

**自动默认值（第一项）：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

enums := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
)
// 第一项（UNKNOWN）自动成为默认值
defaultEnum := enums.GetDefault()
```

**严格的默认值管理：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// 集合必须有默认值
// NewEnums 自动将第一项设为默认值
enums := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
)

// 查找失败返回默认值（永不返回 nil）
notFound := enums.GetByCode(999)  // 返回 UNKNOWN（默认值）
fmt.Printf("回退值: %s\n", notFound.Meta().Desc())  // 无需 nil 检查即可安全使用

// 使用严格模式更改默认值
enums.UnsetDefault()  // 必须先取消设置
enums.SetDefault(enums.MustGetByCode(1))  // 然后设置新默认值

// UnsetDefault 后，查找失败会 panic
// 这强制实施单一使用模式：集合必须有默认值
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/go-xlan/protoenum.svg?variant=adaptive)](https://starchart.cc/go-xlan/protoenum)
