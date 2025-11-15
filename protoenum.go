// Package protoenum: Utilities to handle Protocol Buffer enum metadata management
// Provides type-safe enum descriptors with custom descriptions to enhance enum handling
// Supports generic enum wrapping with description association to improve documentation
//
// protoenum: Protocol Buffer 枚举元数据管理包装工具
// 提供带有自定义描述的类型安全枚举描述符，增强枚举处理能力
// 支持泛型枚举包装和描述关联，提供更好的文档支持
package protoenum

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ProtoEnum establishes the core contract enabling Protocol Buffer enum integration
// Serves as the generic constraint enabling type-safe enum operations across each protobuf enum
// Bridges the native protobuf enum system with enhanced metadata management capabilities
// Important when maintaining compile-time type checking while adding runtime descriptive features
//
// ProtoEnum 建立 Protocol Buffer 枚举集成的基础契约
// 作为泛型约束实现跨所有 protobuf 枚举的类型安全包装操作
// 在原生 protobuf 枚举系统与增强元数据管理能力之间建立桥梁
// 在添加运行时描述特性的同时保持编译时类型安全至关重要
type ProtoEnum interface {
	// String provides the standard name of the enum value as defined in protobuf schema
	// Important when performing serialization, debugging, and human-readable enum identification
	// String 提供 protobuf 模式中定义的枚举值规范名称标识符
	// 在进行序列化、调试和人类可读的枚举识别时至关重要
	String() string
	// Number exposes the underlying numeric wire-format encoding used in protobuf serialization
	// Enables efficient storage, transmission, and support with protobuf specifications
	// Number 暴露 protobuf 序列化中使用的底层数值线格式编码
	// 实现高效存储、传输以及与 protobuf 规范的兼容性
	Number() protoreflect.EnumNumber
}

// Enum wraps a Protocol Buffer enum with extra metadata
// Associates a custom description with the enum value during documentation
// Uses generics to maintain type checking across different enum types
//
// Enum 使用附加元数据包装 Protocol Buffer 枚举
// 在文档化时关联枚举值与自定义描述
// 使用泛型在不同枚举类型间保持类型安全
type Enum[protoEnum ProtoEnum] struct {
	enum        protoEnum // Source Protocol Buffer enum value // 源 Protocol Buffer 枚举值
	description string    // Custom description of the enum // 枚举的自定义描述
}

// NewEnum creates a new Enum instance with the given enum value and description
// Returns a reference to the created Enum instance allowing simple chaining
//
// 使用给定的枚举值和描述创建新的 Enum 包装器
// 返回创建的 Enum 实例指针以便链式调用
func NewEnum[protoEnum ProtoEnum](enum protoEnum, description string) *Enum[protoEnum] {
	return &Enum[protoEnum]{
		enum:        enum,
		description: description,
	}
}

// Base returns the underlying Protocol Buffer enum value
// Provides access to the source enum enabling Protocol Buffer operations
//
// 返回底层的 Protocol Buffer 枚举值
// 提供对源枚举的访问以进行 Protocol Buffer 操作
func (c *Enum[protoEnum]) Base() protoEnum {
	return c.enum
}

// Code returns the numeric code of the enum as int32
// Converts the Protocol Buffer enum value to a standard int32 type
//
// 返回枚举的数字代码作 int32
// 将 Protocol Buffer 枚举数字转换成标准 int32 类型
func (c *Enum[protoEnum]) Code() int32 {
	return int32(c.enum.Number())
}

// Name returns the string name of the enum value
// Gets the Protocol Buffer enum's string representation
//
// 返回枚举值的字符串名称
// 获取 Protocol Buffer 枚举的字符串表示
func (c *Enum[protoEnum]) Name() string {
	return c.enum.String()
}

// Desc returns the custom description of the enum
// Provides human-readable description with documentation purposes
//
// 返回枚举的自定义描述
// 提供人类可读的描述用于文档目的
func (c *Enum[protoEnum]) Desc() string {
	return c.description
}

// Hans returns the Chinese description of the enum
// Alias to GetByDesc method, convenient with Chinese language support
//
// 返回枚举的中文描述
// Desc 方法的别名，方便中文语言支持
func (c *Enum[protoEnum]) Hans() string {
	return c.description
}
