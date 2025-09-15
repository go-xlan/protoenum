[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/protoenum)](https://pkg.go.dev/github.com/go-xlan/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/protoenum/main.svg)](https://coveralls.io/github/go-xlan/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/protoenum.svg)](https://github.com/go-xlan/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/protoenum)](https://goreportcard.com/report/github.com/go-xlan/protoenum)

# protoenum

`protoenum` 是一个 Go 语言包，提供管理 Protobuf 枚举元数据的工具。它将 Protobuf 枚举值与自定义描述包装在一起，并提供枚举集合支持按代码、名称或描述进行快速查找。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **智能枚举管理**：将 Protobuf 枚举与自定义描述和元数据包装
⚡ **多方式查找**：支持通过代码、名称或描述快速查找
🔄 **类型安全操作**：保持 protobuf 类型安全同时增强元数据
🌍 **生产就绪**：经过实战检验的企业级枚举处理方案
📋 **零依赖**：轻量级解决方案，仅使用标准库

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
	"fmt"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
)

// 构建状态枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)

func main() {
	// 从 protobuf 枚举获取增强描述（找不到时返回默认值）
	successStatus := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	fmt.Printf("状态: %s\n", successStatus.Desc())

	// protoenum 与原生枚举间转换（有默认值回退保障）
	statusEnum := enums.GetByName("SUCCESS")
	native := protoenumstatus.StatusEnum(statusEnum.Code())
	fmt.Printf("原生枚举: %v\n", native)

	// 在业务逻辑中使用
	if native == protoenumstatus.StatusEnum_SUCCESS {
		fmt.Println("操作完成！")
	}
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 高级查找方法

```go
package main

import (
	"fmt"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
)

// 构建枚举集合
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, "其它"),
	protoenum.NewEnum(protoenumresult.ResultEnum_PASS, "通过"),
	protoenum.NewEnum(protoenumresult.ResultEnum_FAIL, "出错"),
	protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, "跳过"),
)

func main() {
	// 按枚举代码查找（找不到时返回默认值）
	skipResult := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	fmt.Printf("结果: %s\n", skipResult.Desc())

	// 按枚举名称查找（有默认值回退保障）
	passResult := enums.GetByName("PASS")
	native := protoenumresult.ResultEnum(passResult.Code())
	fmt.Printf("原生: %v\n", native)

	// 使用原生枚举的业务逻辑
	if native == protoenumresult.ResultEnum_PASS {
		fmt.Println("测试通过！")
	}

	// 按中文描述查找（找不到时返回默认值）
	result := enums.GetByDesc("跳过")
	fmt.Printf("名称: %s\n", result.Name())
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)


## API 参考

### 单个枚举操作

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnum(value, desc)` | 创建枚举包装器 | `*Enum[T]` |
| `enum.Code()` | 获取数值代码 | `int32` |
| `enum.Name()` | 获取枚举名称 | `string` |
| `enum.Desc()` | 获取描述 | `string` |

### 集合操作

| 方法 | 说明 | 返回值 |
|------|------|--------|
| `NewEnums(items...)` | 创建枚举集合（第一项作为默认值） | `*Enums[T]` |
| `enums.GetByCode(code)` | 按代码查找（找不到时返回默认值） | `*Enum[T]` |
| `enums.GetByName(name)` | 按名称查找（找不到时返回默认值） | `*Enum[T]` |
| `enums.GetByDesc(desc)` | 按描述查找（找不到时返回默认值） | `*Enum[T]` |
| `enums.SetDefault(enum)` | 动态设置默认值 | `void` |
| `enums.GetDefault()` | 获取当前默认值 | `*Enum[T]` |
| `enums.WithDefaultEnum(enum)` | 链式：通过枚举实例设置默认值 | `*Enums[T]` |
| `enums.WithDefaultCode(code)` | 链式：通过代码设置默认值（找不到则 panic） | `*Enums[T]` |
| `enums.WithDefaultName(name)` | 链式：通过名称设置默认值（找不到则 panic） | `*Enums[T]` |

## 使用示例

### 单个枚举操作

**创建增强枚举包装器：**
```go
statusEnum := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "操作成功")
fmt.Printf("代码: %d, 名称: %s, 描述: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Desc())
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
statusEnums := protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知状态"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)
```

**多种查找方式：**
```go
// 按数字代码查找
if enum := statusEnums.GetByCode(1); enum != nil {
    fmt.Printf("找到: %s\n", enum.Desc())
}

// 按枚举名称查找
if enum := statusEnums.GetByName("SUCCESS"); enum != nil {
    fmt.Printf("状态: %s\n", enum.Desc())
}

// 按中文描述查找
if enum := statusEnums.GetByDesc("成功"); enum != nil {
    fmt.Printf("代码: %d\n", enum.Code())
}
```

### 高级用法


**类型转换模式：**
```go
// 从枚举包装器转换为原生 protobuf 枚举
if statusEnum := enums.GetByName("SUCCESS"); statusEnum != nil {
    native := protoenumstatus.StatusEnum(statusEnum.Code())
    // 在 protobuf 操作中使用原生枚举
}
```

**查找时的错误处理：**
```go
// 安全查找和空值检查
if result := enums.GetByDesc("不存在的描述"); result == nil {
    fmt.Println("未找到枚举")
} else {
    fmt.Printf("找到: %s\n", result.Name())
}
```

### 默认值和链式配置

**自动默认值（第一项）：**
```go
enums := protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
)
// 第一项（UNKNOWN）自动成为默认值
defaultEnum := enums.GetDefault()
```

**链式风格默认值配置：**
```go
// 在初始化时使用链式方法设置默认值
var globalEnums = protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
).WithDefaultCode(0)  // 设置 UNKNOWN 为默认值

// 查找失败时返回默认值而不是 nil
notFound := enums.GetByCode(999)  // 返回默认值（UNKNOWN）而不是 nil
fmt.Printf("回退值: %s\n", notFound.Desc())  // 无需空值检查即可安全使用
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
