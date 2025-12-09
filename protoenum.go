// Package protoenum: Utilities to handle Protocol Buffer enum metadata management
// Provides type-safe enum descriptors with Go native enum binding and custom metadata
// Supports triple generic wrapping: protoEnum, basic, and metaType
// Enables seamless conversion between protobuf enums and Go native enum types
//
// protoenum: Protocol Buffer 枚举元数据管理包装工具
// 提供带有 Go 原生枚举绑定和自定义元数据的类型安全枚举描述符
// 支持三泛型包装：protoEnum、basic 和 metaType
// 实现 protobuf 枚举与 Go 原生枚举类型之间的无缝转换
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

	// comparable constraint matches protobuf enum patterns and enables map-index usage
	// protobuf enums can be compared, this constraint maintains type-safe operations
	// comparable 约束与 protobuf 枚举行为一致，并支持作为 map 键使用
	// Protocol Buffer 枚举本身可比较，此约束确保类型安全
	comparable
}

// Enum wraps a Protocol Buffer enum with Go native enum and custom metadata
// Bridges protobuf enum (protoEnum) with Go native enum (basic) via Basic() method
// Associates custom metadata with the enum value via Meta() method
// Uses triple generics to maintain type checking across protobuf, Go native enum, and metadata
//
// Enum 使用 Go 原生枚举和自定义元数据包装 Protocol Buffer 枚举
// 通过 Basic() 方法桥接 protobuf 枚举 (protoEnum) 和 Go 原生枚举 (basic)
// 通过 Meta() 方法关联枚举值与自定义元数据
// 使用三泛型在 protobuf、Go 原生枚举和元数据类型间保持类型安全
type Enum[protoEnum ProtoEnum, basicEnum comparable, metaType any] struct {
	proto protoEnum // Source Protocol Buffer enum value // 源 Protocol Buffer 枚举值
	basic basicEnum // Go native enum value (e.g. type StatusType string) // Go 原生枚举值（如 type StatusType string）
	meta  metaType  // Custom metadata of the enum // 枚举的自定义元数据
}

// NewEnum creates a new Enum instance binding protobuf enum with Go native enum
// Use this when you just need enum mapping without description
// The basic param accepts Go native enum type (e.g. type StatusType string)
// Returns a reference to the created Enum instance, supporting chained invocation
//
// 创建新的 Enum 实例，绑定 protobuf 枚举与 Go 原生枚举
// 当只需要枚举映射而不需要描述时使用此函数
// basic 参数接受 Go 原生枚举类型（如 type StatusType string）
// 返回创建的 Enum 实例指针以便链式调用
func NewEnum[protoEnum ProtoEnum, basicEnum comparable](proto protoEnum, basic basicEnum) *Enum[protoEnum, basicEnum, *MetaNone] {
	return &Enum[protoEnum, basicEnum, *MetaNone]{
		proto: proto,
		basic: basic,
		meta:  &MetaNone{},
	}
}

// NewEnumWithDesc creates a new Enum instance with protobuf enum, Go native enum, and description
// Use this when you need both enum mapping and human-readable description
// The basic param accepts Go native enum type (e.g. type StatusType string)
// The description param provides custom description used in docs and UI rendering
//
// 创建带有 protobuf 枚举、Go 原生枚举和描述的新 Enum 实例
// 当需要枚举映射和人类可读描述时使用此函数
// basic 参数接受 Go 原生枚举类型（如 type StatusType string）
// description 参数提供用于文档和显示的自定义描述
func NewEnumWithDesc[protoEnum ProtoEnum, basicEnum comparable](proto protoEnum, basic basicEnum, description string) *Enum[protoEnum, basicEnum, *MetaDesc] {
	return &Enum[protoEnum, basicEnum, *MetaDesc]{
		proto: proto,
		basic: basic,
		meta:  &MetaDesc{description: description},
	}
}

// NewEnumWithMeta creates a new Enum instance with protobuf enum, Go native enum, and custom metadata
// Use this when you need customized metadata types beyond simple string description
// The basic param accepts Go native enum type (e.g. type StatusType string)
// The meta param accepts custom metadata types (e.g. i18n descriptions with multiple languages)
//
// 创建带有 protobuf 枚举、Go 原生枚举和自定义元数据的新 Enum 实例
// 当需要超越简单字符串描述的灵活元数据类型时使用此函数
// basic 参数接受 Go 原生枚举类型（如 type StatusType string）
// meta 参数接受任意自定义元数据类型（如双语描述）
func NewEnumWithMeta[protoEnum ProtoEnum, basicEnum comparable, metaType any](proto protoEnum, basic basicEnum, meta metaType) *Enum[protoEnum, basicEnum, metaType] {
	return &Enum[protoEnum, basicEnum, metaType]{
		proto: proto,
		basic: basic,
		meta:  meta,
	}
}

// Proto returns the underlying Protocol Buffer enum value
// Provides access to the source enum enabling Protocol Buffer operations
//
// 返回底层的 Protocol Buffer 枚举值
// 提供对源枚举的访问以进行 Protocol Buffer 操作
func (c *Enum[protoEnum, basicEnum, metaType]) Proto() protoEnum {
	return c.proto
}

// Basic returns the Go native enum value associated with this enum
// Enables type-safe conversion from protobuf enum to Go native enum (e.g. type StatusType string)
// Use this to get the basic enum value when working with Go native enum patterns
// Bridges protobuf enums with existing Go enum-based business logic with ease
//
// 返回与此枚举关联的 Go 原生枚举值
// 实现从 protobuf 枚举到 Go 原生枚举的类型安全转换（如 type StatusType string）
// 在使用 Go 原生枚举模式时使用此方法获取 basic 枚举值
// 无缝桥接 protobuf 枚举与现有基于 Go 枚举的业务逻辑
func (c *Enum[protoEnum, basicEnum, metaType]) Basic() basicEnum {
	return c.basic
}

// Code returns the numeric code of the enum as int32
// Converts the Protocol Buffer enum value to a standard int32 type
//
// 返回枚举的数字代码作 int32
// 将 Protocol Buffer 枚举数字转换成标准 int32 类型
func (c *Enum[protoEnum, basicEnum, metaType]) Code() int32 {
	return int32(c.proto.Number())
}

// Name returns the string name of the enum value
// Gets the Protocol Buffer enum's string representation
//
// 返回枚举值的字符串名称
// 获取 Protocol Buffer 枚举的字符串表示
func (c *Enum[protoEnum, basicEnum, metaType]) Name() string {
	return c.proto.String()
}

// Meta returns the metadata associated with this enum
// Provides access to custom metadata like description via MetaDesc
// Use this when you need to access extended enum metadata
//
// 返回与此枚举关联的元数据
// 提供对自定义元数据（如通过 MetaDesc 获取描述）的访问
// 在需要访问额外的枚举元数据时使用此方法
func (c *Enum[protoEnum, basicEnum, metaType]) Meta() metaType {
	return c.meta
}
