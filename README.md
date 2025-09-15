[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/protoenum/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/protoenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/protoenum)](https://pkg.go.dev/github.com/go-xlan/protoenum)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/protoenum/main.svg)](https://coveralls.io/github/go-xlan/protoenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/protoenum.svg)](https://github.com/go-xlan/protoenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/protoenum)](https://goreportcard.com/report/github.com/go-xlan/protoenum)

# protoenum

`protoenum` provides utilities for managing Protobuf enum metadata in Go. It wraps Protobuf enum values with custom descriptions and offers enum collections for easy lookup by code, name, or description.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Core Features

🎯 **Smart Enum Management**: Wrap Protobuf enums with custom descriptions and metadata
⚡ **Multi-Lookup Support**: Fast lookup by code, name, or description
🔄 **Type-Safe Operations**: Preserve protobuf type safety with enhanced metadata
🌍 **Production Ready**: Battle-tested enum handling for enterprise applications
📋 **Zero Dependencies**: Lightweight solution with standard library

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
	"fmt"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
)

// Build status enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)

func main() {
	// Get enhanced description from protobuf enum (returns default if not found)
	successStatus := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	fmt.Printf("Status: %s\n", successStatus.Desc())

	// Convert between protoenum and native enum (safe with default fallback)
	statusEnum := enums.GetByName("SUCCESS")
	native := protoenumstatus.StatusEnum(statusEnum.Code())
	fmt.Printf("Native enum: %v\n", native)

	// Use in business logic
	if native == protoenumstatus.StatusEnum_SUCCESS {
		fmt.Println("Operation completed!")
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Advanced Lookup Methods

```go
package main

import (
	"fmt"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
)

// Build enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, "其它"),
	protoenum.NewEnum(protoenumresult.ResultEnum_PASS, "通过"),
	protoenum.NewEnum(protoenumresult.ResultEnum_FAIL, "出错"),
	protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, "跳过"),
)

func main() {
	// Lookup by enum code (returns default if not found)
	skipResult := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	fmt.Printf("Result: %s\n", skipResult.Desc())

	// Lookup by enum name (safe with default fallback)
	passResult := enums.GetByName("PASS")
	native := protoenumresult.ResultEnum(passResult.Code())
	fmt.Printf("Native: %v\n", native)

	// Business logic with native enum
	if native == protoenumresult.ResultEnum_PASS {
		fmt.Println("Test passed!")
	}

	// Lookup by Chinese description (returns default if not found)
	result := enums.GetByDesc("跳过")
	fmt.Printf("Name: %s\n", result.Name())
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)


## API Reference

### Individual Enum Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnum(value, desc)` | Create enum wrapper | `*Enum[T]` |
| `enum.Code()` | Get numeric code | `int32` |
| `enum.Name()` | Get enum name | `string` |
| `enum.Desc()` | Get description | `string` |

### Collection Operations

| Method | Description | Returns |
|--------|-------------|--------|
| `NewEnums(items...)` | Create enum collection (first item as default) | `*Enums[T]` |
| `enums.GetByCode(code)` | Lookup by code (returns default if not found) | `*Enum[T]` |
| `enums.GetByName(name)` | Lookup by name (returns default if not found) | `*Enum[T]` |
| `enums.GetByDesc(desc)` | Lookup by description (returns default if not found) | `*Enum[T]` |
| `enums.SetDefault(enum)` | Set default value dynamically | `void` |
| `enums.GetDefault()` | Get current default value | `*Enum[T]` |
| `enums.WithDefaultEnum(enum)` | Chain: set default by enum instance | `*Enums[T]` |
| `enums.WithDefaultCode(code)` | Chain: set default by code (panics if not found) | `*Enums[T]` |
| `enums.WithDefaultName(name)` | Chain: set default by name (panics if not found) | `*Enums[T]` |

## Examples

### Working with Individual Enums

**Creating enhanced enum wrapper:**
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
// By numeric code
if enum := statusEnums.GetByCode(1); enum != nil {
    fmt.Printf("Found: %s\n", enum.Desc())
}

// By enum name
if enum := statusEnums.GetByName("SUCCESS"); enum != nil {
    fmt.Printf("Status: %s\n", enum.Desc())
}

// By Chinese description
if enum := statusEnums.GetByDesc("成功"); enum != nil {
    fmt.Printf("Code: %d\n", enum.Code())
}
```

### Advanced Usage


**Type conversion patterns:**
```go
// Convert from enum wrapper to native protobuf enum
if statusEnum := enums.GetByName("SUCCESS"); statusEnum != nil {
    native := protoenumstatus.StatusEnum(statusEnum.Code())
    // Use native enum in protobuf operations
}
```

**Error handling with lookups:**
```go
// Safe lookup with nil check
if result := enums.GetByDesc("不存在的描述"); result == nil {
    fmt.Println("Enum not found")
} else {
    fmt.Printf("Found: %s\n", result.Name())
}
```

### Default Values and Chain Configuration

**Automatic default value (first item):**
```go
enums := protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
)
// First item (UNKNOWN) automatically becomes default
defaultEnum := enums.GetDefault()
```

**Chain-style default configuration:**
```go
// Set default using chain methods during initialization
var globalEnums = protoenum.NewEnums(
    protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
    protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
).WithDefaultCode(0)  // Set UNKNOWN as default

// Lookup failures return default instead of nil
notFound := enums.GetByCode(999)  // Returns default (UNKNOWN) instead of nil
fmt.Printf("Fallback: %s\n", notFound.Desc())  // Safe to use without nil check
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
