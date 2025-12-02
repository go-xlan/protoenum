[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/protoenum)](https://pkg.go.dev/github.com/go-xlan/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/protoenum/main.svg)](https://coveralls.io/github/go-xlan/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.23--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/protoenum.svg)](https://github.com/go-xlan/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/protoenum)](https://goreportcard.com/report/github.com/go-xlan/protoenum)

# protoenum

`protoenum` provides utilities to manage Protobuf enum metadata in Go. It bridges Protobuf enums with Go native enums (`type StatusType string`) via the `Pure()` method, and offers enum collections with simple code, name, and pure-value lookups.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Core Features

🎯 **Smart Enum Management**: Wrap Protobuf enums with Go native enums and custom metadata
🔗 **Go Native Enum Bridge**: Seamless conversion via `Pure()` method to Go native enum types
⚡ **Multi-Lookup Support**: Fast code, name, and pure-value lookups
🔄 **Type-Safe Operations**: Triple generics preserve type safety across protobuf, Go native enums, and metadata
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

// StatusType represents a Go native enum of status
type StatusType string

const (
	StatusTypeUnknown StatusType = "unknown"
	StatusTypeSuccess StatusType = "success"
	StatusTypeFailure StatusType = "failure"
)

// Build status enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
)

func main() {
	// Get Go native enum from protobuf enum (returns default when not found)
	item := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	zaplog.LOG.Debug("pure", zap.String("msg", string(item.Pure())))

	// Convert between protoenum and native enum (safe with default fallback)
	enum := enums.GetByName("SUCCESS")
	base := protoenumstatus.StatusEnum(enum.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Use in business logic
	if base == protoenumstatus.StatusEnum_SUCCESS {
		zaplog.LOG.Debug("done")
	}

	// Get default plain enum value (first item becomes default)
	defaultPure := enums.GetDefaultPure()
	zaplog.LOG.Debug("default", zap.String("msg", string(defaultPure)))
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

// ResultType represents a Go native enum of result
type ResultType string

const (
	ResultTypeUnknown ResultType = "unknown"
	ResultTypePass    ResultType = "pass"
	ResultTypeMiss    ResultType = "miss"
	ResultTypeSkip    ResultType = "skip"
)

// Build enum collection with description
var enums = protoenum.NewEnums(
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown, "其它"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_PASS, ResultTypePass, "通过"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_MISS, ResultTypeMiss, "出错"),
	protoenum.NewEnumWithDesc(protoenumresult.ResultEnum_SKIP, ResultTypeSkip, "跳过"),
)

func main() {
	// Lookup using enum code (returns default when not found)
	skip := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	zaplog.LOG.Debug("pure", zap.String("msg", string(skip.Pure())))
	zaplog.LOG.Debug("desc", zap.String("msg", skip.Meta().Desc()))

	// Lookup using Go native enum value (type-safe)
	pass := enums.GetByPure(ResultTypePass)
	base := protoenumresult.ResultEnum(pass.Code())
	zaplog.LOG.Debug("base", zap.String("msg", base.String()))

	// Business logic with native enum
	if base == protoenumresult.ResultEnum_PASS {
		zaplog.LOG.Debug("pass")
	}

	// Lookup using enum name (safe with default fallback)
	miss := enums.GetByName("MISS")
	zaplog.LOG.Debug("pure", zap.String("msg", string(miss.Pure())))
	zaplog.LOG.Debug("desc", zap.String("msg", miss.Meta().Desc()))

	// List each plain enum value in defined sequence
	pures := enums.ListPures()
	for _, pure := range pures {
		zaplog.LOG.Debug("list", zap.String("pure", string(pure)))
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)


## API Reference

### Single Enum Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnum(protoEnum, plainEnum)` | Create enum instance without metadata | `*Enum[P, E, *MetaNone]` |
| `NewEnumWithDesc(protoEnum, plainEnum, desc)` | Create enum instance with description | `*Enum[P, E, *MetaDesc]` |
| `NewEnumWithMeta(protoEnum, plainEnum, meta)` | Create enum instance with custom metadata | `*Enum[P, E, M]` |
| `enum.Base()` | Get underlying protobuf enum | `P` |
| `enum.Code()` | Get numeric code | `int32` |
| `enum.Name()` | Get enum name | `string` |
| `enum.Pure()` | Get Go native enum value | `E` |
| `enum.Meta()` | Get custom metadata | `M` |

### Collection Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnums(items...)` | Create collection with strict validation (first item becomes default) | `*Enums[P, E, M]` |
| `enums.GetByEnum(enum)` | Lookup by protobuf enum (returns default if not found, panics if no default) | `*Enum[P, E, M]` |
| `enums.GetByCode(code)` | Lookup by code (returns default if not found, panics if no default) | `*Enum[P, E, M]` |
| `enums.GetByName(name)` | Lookup by name (returns default if not found, panics if no default) | `*Enum[P, E, M]` |
| `enums.GetByPure(pure)` | Lookup by Go native enum (returns default if not found, panics if no default) | `*Enum[P, E, M]` |
| `enums.MustGetByEnum(enum)` | Strict lookup by protobuf enum (panics if not found) | `*Enum[P, E, M]` |
| `enums.MustGetByCode(code)` | Strict lookup by code (panics if not found) | `*Enum[P, E, M]` |
| `enums.MustGetByName(name)` | Strict lookup by name (panics if not found) | `*Enum[P, E, M]` |
| `enums.MustGetByPure(pure)` | Strict lookup by Go native enum (panics if not found) | `*Enum[P, E, M]` |
| `enums.ListEnums()` | Returns a slice of each protoEnum value | `[]P` |
| `enums.ListPures()` | Returns a slice of each plainEnum value | `[]E` |
| `enums.ListValidEnums()` | Returns protoEnum values excluding default | `[]P` |
| `enums.ListValidPures()` | Returns plainEnum values excluding default | `[]E` |
| `enums.GetDefault()` | Get current default value (panics if unset) | `*Enum[P, E, M]` |
| `enums.GetDefaultEnum()` | Get default protoEnum value (panics if unset) | `P` |
| `enums.GetDefaultPure()` | Get default plainEnum value (panics if unset) | `E` |
| `enums.SetDefault(enum)` | Set default (requires no existing default) | `void` |
| `enums.UnsetDefault()` | Remove default (requires existing default) | `void` |
| `enums.WithDefaultEnum(enum)` | Chain: set default by enum instance | `*Enums[P, E, M]` |
| `enums.WithDefaultCode(code)` | Chain: set default by code (panics if not found) | `*Enums[P, E, M]` |
| `enums.WithDefaultName(name)` | Chain: set default by name (panics if not found) | `*Enums[P, E, M]` |
| `enums.WithUnsetDefault()` | Chain: remove default value | `*Enums[P, E, M]` |

## Examples

### Working with Single Enums

**Creating enhanced enum instance:**
```go
type StatusType string
const StatusTypeSuccess StatusType = "success"

statusEnum := protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "操作成功")
fmt.Printf("Code: %d, Name: %s, Pure: %s, Description: %s\n",
    statusEnum.Code(), statusEnum.Name(), statusEnum.Pure(), statusEnum.Meta().Desc())
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

**Multiple lookup methods:**
```go
// Using numeric code - returns valid enum (default if not found)
enum := statusEnums.GetByCode(1)
fmt.Printf("Found: %s\n", enum.Meta().Desc())

// Using enum name - guaranteed non-nil
enum = statusEnums.GetByName("SUCCESS")
fmt.Printf("Status: %s\n", enum.Meta().Desc())

// Using Go native enum value - type-safe lookup
enum = statusEnums.GetByPure(StatusTypeSuccess)
fmt.Printf("Pure: %s\n", enum.Pure())

// Strict lookup - panics if not found (no default fallback)
enum = statusEnums.MustGetByCode(1)
fmt.Printf("Strict: %s\n", enum.Meta().Desc())
```

**Listing values:**
```go
// Get a slice of each registered proto enum
protoEnums := statusEnums.ListEnums()
// > [UNKNOWN, SUCCESS, FAILURE]

// Get a slice of each registered plain Go enum
plainEnums := statusEnums.ListPures()
// > ["unknown", "success", "failure"]

// Get valid values (excluding default)
validEnums := statusEnums.ListValidEnums()
// > [SUCCESS, FAILURE] (UNKNOWN is default, excluded)

validPures := statusEnums.ListValidPures()
// > ["success", "failure"]
```

### Advanced Usage

**Go native enum bridge via Pure():**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// Bridge protobuf enum to Go native enum
enum := enums.GetByCode(1)
pureValue := enum.Pure()  // Returns StatusType("success")

// Use in business logic with Go native enum
switch pureValue {
case StatusTypeSuccess:
    fmt.Println("Operation succeeded")
case StatusTypeUnknown:
    fmt.Println("Unknown status")
}

// Lookup using Go native enum value
found := enums.GetByPure(StatusTypeSuccess)
fmt.Printf("Code: %d, Name: %s\n", found.Code(), found.Name())
```

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
result := enums.MustGetByCode(1)  // Panics if code not in collection
fmt.Printf("Found: %s\n", result.Name())

// GetByXxx returns default on unknown values
result = enums.GetByCode(999)  // Returns default (UNKNOWN)
fmt.Printf("Fallback: %s\n", result.Name())
```

### Default Values and Chain Configuration

**Automatic default value (first item):**
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
// First item (UNKNOWN) becomes default on creation
defaultEnum := enums.GetDefault()
```

**Strict default management:**
```go
type StatusType string
const (
    StatusTypeUnknown StatusType = "unknown"
    StatusTypeSuccess StatusType = "success"
)

// Collections MUST have a default value
// NewEnums sets first item as default on init
enums := protoenum.NewEnums(
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown, "未知"),
    protoenum.NewEnumWithDesc(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess, "成功"),
)

// Lookup failures return default (not nil)
notFound := enums.GetByCode(999)  // Returns UNKNOWN (default)
fmt.Printf("Fallback: %s\n", notFound.Meta().Desc())  // Safe without nil check

// Change default using strict pattern
enums.UnsetDefault()  // Must unset first
enums.SetDefault(enums.MustGetByCode(1))  // Then set new default

// Once UnsetDefault called, lookups panic if not found
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
