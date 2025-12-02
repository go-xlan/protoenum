[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/protoenum)](https://pkg.go.dev/github.com/go-xlan/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/protoenum/main.svg)](https://coveralls.io/github/go-xlan/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.23--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/protoenum.svg)](https://github.com/go-xlan/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/protoenum)](https://goreportcard.com/report/github.com/go-xlan/protoenum)

# protoenum

`protoenum` 是一个 Go 语言包，提供管理 Protobuf 枚举元数据的工具。它通过 `Pure()` 方法桥接 Protobuf 枚举和 Go 原生枚举（`type StatusType string`），并提供枚举集合支持简单的代码、名称和 Pure 值查找。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **智能枚举管理**：将 Protobuf 枚举与 Go 原生枚举和自定义元数据包装
🔗 **Go 原生枚举桥接**：通过 `Pure()` 方法无缝转换到 Go 原生枚举类型
⚡ **多方式查找**：支持代码、名称和 Pure 值快速查找
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

// StatusType 代表状态的 Go 原生枚举
type StatusType string

const (
	StatusTypeUnknown StatusType = "unknown"
	StatusTypeSuccess StatusType = "success"
	StatusTypeFailure StatusType = "failure"
)

// 构建状态枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
)

func main() {
	// 从 protobuf 枚举获取 Go 原生枚举（找不到时返回默认值）
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("pure", zap.String("msg", string(item.Pure())))

	// 在 protoenum 和原生枚举之间转换（安全且有默认值回退）
	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// 在业务逻辑中使用
	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}
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

// ResultType 代表结果的 Go 原生枚举
type ResultType string

const (
	ResultTypeUnknown ResultType = "unknown"
	ResultTypePass    ResultType = "pass"
	ResultTypeMiss    ResultType = "miss"
	ResultTypeSkip    ResultType = "skip"
)

// 构建带描述的枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown, "其它"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_PASS, ResultTypePass, "通过"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_MISS, ResultTypeMiss, "出错"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_SKIP, ResultTypeSkip, "跳过"),
)

func main() {
	// 按枚举代码查找（找不到时返回默认值）
	skip := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	zaplog.LOG.Debug("pure", zap.String("msg", string(skip.Pure())))
	zaplog.LOG.Debug("desc", zap.String("msg", skip.Meta().Desc()))

	// 按 Go 原生枚举值查找（类型安全查找）
	pass := enums.GetByPure(ResultTypePass)
	base := protoenumresult.ResultEnum(pass.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// 使用原生枚举的业务逻辑
	if base == protoenumresult.ResultEnum_PASS {
		zaplog.LOG.Debug("pass")
	}

	// 按枚举名称查找（安全且有默认值回退）
	miss := enums.GetByName("MISS")
	zaplog.LOG.Debug("pure", zap.String("msg", string(miss.Pure())))
	zaplog.LOG.Debug("desc", zap.String("msg", miss.Meta().Desc()))
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)


## API 参考

### 单个枚举操作

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnum(protoEnum, plainEnum)` | 创建枚举实例（无元数据） | `*Enum[P, E, *MetaNone]` |
| `NewEnumWithDesc(protoEnum, plainEnum, desc)` | 创建枚举实例（带描述） | `*Enum[P, E, *MetaDesc]` |
| `NewEnumWithMeta(protoEnum, plainEnum, meta)` | 创建枚举实例（带自定义元数据） | `*Enum[P, E, M]` |
| `enum.Base()` | 获取底层 protobuf 枚举 | `P` |
| `enum.Code()` | 获取数值代码 | `int32` |
| `enum.Name()` | 获取枚举名称 | `string` |
| `enum.Pure()` | 获取 Go 原生枚举值 | `E` |
| `enum.Meta()` | 获取自定义元数据 | `M` |

### 集合操作

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnums(items...)` | 创建集合并严格验证（第一项成为默认值） | `*Enums[P, E, M]` |
| `enums.GetByEnum(enum)` | 按 protobuf 枚举查找（找不到返回默认值，无默认值则 panic） | `*Enum[P, E, M]` |
| `enums.GetByCode(code)` | 按代码查找（找不到返回默认值，无默认值则 panic） | `*Enum[P, E, M]` |
| `enums.GetByName(name)` | 按名称查找（找不到返回默认值，无默认值则 panic） | `*Enum[P, E, M]` |
| `enums.GetByPure(pure)` | 按 Go 原生枚举查找（找不到返回默认值，无默认值则 panic） | `*Enum[P, E, M]` |
| `enums.MustGetByEnum(enum)` | 严格按 protobuf 枚举查找（找不到则 panic） | `*Enum[P, E, M]` |
| `enums.MustGetByCode(code)` | 严格按代码查找（找不到则 panic） | `*Enum[P, E, M]` |
| `enums.MustGetByName(name)` | 严格按名称查找（找不到则 panic） | `*Enum[P, E, M]` |
| `enums.MustGetByPure(pure)` | 严格按 Go 原生枚举查找（找不到则 panic） | `*Enum[P, E, M]` |
| `enums.ListEnums()` | 返回所有 protoEnum 值的切片 | `[]P` |
| `enums.ListPures()` | 返回所有 plainEnum 值的切片 | `[]E` |
| `enums.GetDefault()` | 获取当前默认值（未设置则 panic） | `*Enum[P, E, M]` |
| `enums.SetDefault(enum)` | 设置默认值（要求当前无默认值） | `void` |
| `enums.UnsetDefault()` | 移除默认值（要求当前有默认值） | `void` |
| `enums.WithDefaultEnum(enum)` | 链式：通过枚举实例设置默认值 | `*Enums[P, E, M]` |
| `enums.WithDefaultCode(code)` | 链式：通过代码设置默认值（找不到则 panic） | `*Enums[P, E, M]` |
| `enums.WithDefaultName(name)` | 链式：通过名称设置默认值（找不到则 panic） | `*Enums[P, E, M]` |
| `enums.WithUnsetDefault()` | 链式：移除默认值 | `*Enums[P, E, M]` |

## 使用示例

### 单个枚举操作

**创建增强枚举包装器：**
```go
type StatusType string
const StatusTypeSuccess StatusType = "success"

statusEnum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "操作成功")
fmt.Printf("代码: %d, 名称: %s, Pure: %s, 描述: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Pure(), statusEnum.Meta().Desc())
```

**访问底层 protobuf 枚举：**
```go
originalEnum := statusEnum.Base()
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
enum = statusEnums.GetByPure(StatusTypeSuccess)
fmt.Printf("Pure: %s\n", enum.Pure())

// 严格按 Go 原生枚举值查找（找不到则 panic）
enum = statusEnums.MustGetByCode(1)
fmt.Printf("严格: %s\n", enum.Meta().Desc())
```

**列出所有值:**
```go
// 获取所有已注册的 proto 枚举切片
allProtoEnums := statusEnums.ListEnums()
// > [UNKNOWN, SUCCESS, FAILURE]

// 获取所有已注册的 Go 原生枚举切片
allPlainEnums := statusEnums.ListPures()
// > ["unknown", "success", "failure"]
```

### 高级用法

**通过 Pure() 桥接 Go 原生枚举：**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// 桥接 protobuf 枚举到 Go 原生枚举
enum := enums.GetByCode(1)
pureValue := enum.Pure()  // 返回 StatusType("success")

// 在业务逻辑中使用 Go 原生枚举
switch pureValue {
case StatusTypeSuccess:
    fmt.Println("操作成功")
case StatusTypeUnknown:
    fmt.Println("未知状态")
}

// 通过 Go 原生枚举值查找
found := enums.GetByPure(StatusTypeSuccess)
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
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
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
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包编程快乐！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/go-xlan/protoenum.svg?variant=adaptive)](https://starchart.cc/go-xlan/protoenum)
