[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/protoenum)](https://pkg.go.dev/github.com/go-xlan/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/protoenum/main.svg)](https://coveralls.io/github/go-xlan/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/go-xlan/protoenum)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/protoenum.svg)](https://github.com/go-xlan/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/protoenum)](https://goreportcard.com/report/github.com/go-xlan/protoenum)

# protoenum

`protoenum` provides utilities to manage Protobuf enum metadata in Go. It wraps Protobuf enum values with custom descriptions and offers enum collections with simple lookup by code, name, or description.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Core Features

🎯 **Smart Enum Management**: Wrap Protobuf enums with custom descriptions and metadata
⚡ **Multi-Lookup Support**: Fast lookup by code, name, or description with strict validation
🔄 **Type-Safe Operations**: Preserve protobuf type safe operations with enhanced metadata
🛡️ **Strict Design**: Single usage pattern prevents misuse with required defaults
🌍 **Production Grade**: Battle-tested enum handling in enterprise applications

## Installation

```bash
go get github.com/go-xlan/protoenum
```

## Quick Start

### Define Proto Enum

The project includes example proto files:
- `protoenumstatus.proto` - Basic status enum
- `protoenumresult.proto` - Test result enum

### Basic Collection Usage

```go
package main

import (
	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Build status enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)

func main() {
	// Get enhanced description from protobuf enum (returns default when not found)
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("desc", zap.String("msg", item.Desc()))

	// Convert between protoenum and native enum (safe with default fallback)
	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Use in business logic
	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Advanced Lookup Methods

```go
package main

import (
	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Build enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, "其它"),
	protoenum.NewEnum(protoenumresult.ResultEnum_PASS, "通过"),
	protoenum.NewEnum(protoenumresult.ResultEnum_FAIL, "出错"),
	protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, "跳过"),
)

func main() {
	// Lookup by enum code (returns default when not found)
	skip := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	zaplog.LOG.Debug("desc", zap.String("msg", skip.Desc()))

	// Lookup by enum name (safe with default fallback)
	pass := enums.GetByName("PASS")
	base := protoenumresult.ResultEnum(pass.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Business logic with native enum
	if base == protoenumresult.ResultEnum_PASS {
		zaplog.LOG.Debug("pass")
	}

	// Lookup by Chinese description (returns default when not found)
	skip = enums.GetByDesc("跳过")
	zaplog.LOG.Debug("name", zap.String("msg", skip.Name()))
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)


## API Reference

### Single Enum Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnum(value, desc)` | Create enum instance | `*Enum[T]` |
| `enum.Base()` | Get underlying protobuf enum | `T` |
| `enum.Code()` | Get numeric code | `int32` |
| `enum.Name()` | Get enum name | `string` |
| `enum.Desc()` | Get description | `string` |
| `enum.Hans()` | Get Chinese description (alias to Desc) | `string` |

### Collection Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnums(items...)` | Create collection with strict validation (first item becomes default) | `*Enums[T]` |
| `enums.GetByEnum(enum)` | Lookup by protobuf enum (returns default if not found, panics if no default) | `*Enum[T]` |
| `enums.GetByCode(code)` | Lookup by code (returns default if not found, panics if no default) | `*Enum[T]` |
| `enums.GetByName(name)` | Lookup by name (returns default if not found, panics if no default) | `*Enum[T]` |
| `enums.GetByDesc(desc)` | Lookup by description (returns default if not found, panics if no default) | `*Enum[T]` |
| `enums.GetByHans(hans)` | Lookup by Chinese description (alias to GetByDesc) | `*Enum[T]` |
| `enums.MustGetByEnum(enum)` | Strict lookup by protobuf enum (panics if not found) | `*Enum[T]` |
| `enums.MustGetByCode(code)` | Strict lookup by code (panics if not found) | `*Enum[T]` |
| `enums.MustGetByName(name)` | Strict lookup by name (panics if not found) | `*Enum[T]` |
| `enums.MustGetByDesc(desc)` | Strict lookup by description (panics if not found) | `*Enum[T]` |
| `enums.MustGetByHans(hans)` | Strict lookup by Chinese description (alias to MustGetByDesc) | `*Enum[T]` |
| `enums.GetDefault()` | Get current default value (panics if unset) | `*Enum[T]` |
| `enums.SetDefault(enum)` | Set default (requires no existing default) | `void` |
| `enums.UnsetDefault()` | Remove default (requires existing default) | `void` |
| `enums.WithDefaultEnum(enum)` | Chain: set default by enum instance | `*Enums[T]` |
| `enums.WithDefaultCode(code)` | Chain: set default by code (panics if not found) | `*Enums[T]` |
| `enums.WithDefaultName(name)` | Chain: set default by name (panics if not found) | `*Enums[T]` |
| `enums.WithUnsetDefault()` | Chain: remove default value | `*Enums[T]` |

## Examples

### Working with Single Enums

**Creating enhanced enum instance:**
```go
statusEnum := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "操作成功")
fmt.Printf("Code: %d, Name: %s, Description: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Desc())
```

**Accessing underlying protobuf enum:**
```go
originalEnum := statusEnum.Base()
if originalEnum == protoenumstatus.StatusEnum_SUCCESS {
    fmt.Println("Success status detected")
}
```

### Collection Operations

**Building enum collections:**
```go
statusEnums := protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知状态"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)
```

**Multiple lookup methods:**
```go
// By numeric code - always returns valid enum (default if not found)
enum := statusEnums.GetByCode(1)
fmt.Printf("Found: %s\n", enum.Desc())

// By enum name - guaranteed non-nil
enum = statusEnums.GetByName("SUCCESS")
fmt.Printf("Status: %s\n", enum.Desc())

// By Chinese description - safe with default fallback
enum = statusEnums.GetByDesc("成功")
fmt.Printf("Code: %d\n", enum.Code())

// Strict lookup - panics if not found (no default fallback)
enum = statusEnums.MustGetByCode(1)
fmt.Printf("Strict: %s\n", enum.Desc())
```

### Advanced Usage


**Type conversion patterns:**
```go
// Convert from enum instance to native protobuf enum
// Always returns valid enum (with default fallback)
statusEnum := enums.GetByName("SUCCESS")
native := protoenumstatus.StatusEnum(statusEnum.Code())
// Use native enum in protobuf operations with safe access
```

**Strict validation patterns:**
```go
// Use MustGetByXxx with strict validation (panics if not found)
result := enums.MustGetByDesc("成功")  // Panics if description not in collection
fmt.Printf("Found: %s\n", result.Name())

// GetByXxx returns default on unknown values
result = enums.GetByDesc("不存在的描述")  // Returns default (UNKNOWN)
fmt.Printf("Fallback: %s\n", result.Name())
```

### Default Values and Chain Configuration

**Automatic default value (first item):**
```go
enums := protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
)
// First item (UNKNOWN) becomes default on creation
defaultEnum := enums.GetDefault()
```

**Strict default management:**
```go
// Collections MUST have a default value
// NewEnums sets first item as default on init
enums := protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
)

// Lookup failures return default (not nil)
notFound := enums.GetByCode(999)  // Returns UNKNOWN (default)
fmt.Printf("Fallback: %s\n", notFound.Desc())  // Safe without nil check

// Change default using strict pattern
enums.UnsetDefault()  // Must unset first
enums.SetDefault(enums.MustGetByCode(1))  // Then set new default

// After UnsetDefault, lookups panic if not found
// This enforces single usage pattern: collections must have defaults
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-xlan/protoenum.svg?variant=adaptive)](https://starchart.cc/go-xlan/protoenum)
