// Package protoenum: Collection management to handle Protocol Buffer enum metadata
// Provides indexed collections of enum descriptors with multiple lookup methods
// Enables fast lookup using code, name, and basic value with efficient enum handling
//
// protoenum: Protocol Buffer 枚举元数据集合管理
// 提供带有多种查找方法的枚举描述符索引集合
// 支持按代码、名称或 basic 枚举值快速检索，实现高效枚举处理
package protoenum

import (
	"slices"

	"github.com/go-xlan/protoenum/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/tern/slicetern"
)

// Enums manages a collection of Enum instances with indexed lookups
// Maintains multiple maps enabling efficient lookup using different identifiers
// Provides O(1) lookup when searching proto, code, name, and basic value
// Includes a configurable default value returned when lookups miss
//
// Enums 管理 Enum 实例集合并提供索引查找
// 维护四个映射表以通过不同标识符高效检索
// 为 proto、代码、名称和 basic 枚举值搜索提供 O(1) 查找性能
// 支持在查找失败时返回可选的默认值
type Enums[P ProtoEnum, B comparable, M any] struct {
	enumElements []*Enum[P, B, M]          // Holds complete Enum instances in defined sequence // 存放所有 Enum 实例，并维持其定义的次序
	mapProtoEnum map[P]*Enum[P, B, M]      // Map from proto enum to Enum // 从 proto 枚举到 Enum 的映射
	mapCode2Enum map[int32]*Enum[P, B, M]  // Map from numeric code to Enum // 从数字代码到 Enum 的映射
	mapName2Enum map[string]*Enum[P, B, M] // Map from name string to Enum // 从名称字符串到 Enum 的映射
	mapBasicEnum map[B]*Enum[P, B, M]      // Map from basic enum to Enum // 从 basic 枚举到 Enum 的映射
	defaultValue *Enum[P, B, M]            // Configurable default value when lookup misses // 查找失败时的可选默认值
	defaultValid *bool                     // When true, default is treated as valid in ListValidXxx // 为 true 时，ListValidXxx 将默认值视为有效
}

// NewEnums creates a new Enums collection from the given Enum instances
// Builds indexed maps enabling efficient lookup using proto, code, name, and basic value
// The first item becomes the default value if provided
// Returns a reference to the created Enums collection, usable in lookup operations
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 构建索引映射以通过 proto、代码、名称和 basic 枚举值高效查找
// 如果提供了参数，第一个项成为默认值
// 返回创建的 Enums 集合指针，可用于各种查找操作
func NewEnums[P ProtoEnum, B comparable, M any](params ...*Enum[P, B, M]) *Enums[P, B, M] {
	res := &Enums[P, B, M]{
		enumElements: slices.Clone(params), // Clone the slice to preserve the defined sequence of enum elements // 克隆切片以保持枚举元素的定义次序
		mapProtoEnum: make(map[P]*Enum[P, B, M], len(params)),
		mapCode2Enum: make(map[int32]*Enum[P, B, M], len(params)),
		mapName2Enum: make(map[string]*Enum[P, B, M], len(params)),
		mapBasicEnum: make(map[B]*Enum[P, B, M], len(params)),
		defaultValue: slicetern.V0(params), // Set first item as default if available // 如果有参数，将第一个设置为默认值
		defaultValid: nil,
	}
	for _, enum := range params {
		must.Full(enum)

		// Check proto collision // 检查 proto 枚举冲突
		must.Null(res.mapProtoEnum[enum.Proto()])
		res.mapProtoEnum[enum.Proto()] = enum
		// Check code collision // 检查代码冲突
		must.Null(res.mapCode2Enum[enum.Code()])
		res.mapCode2Enum[enum.Code()] = enum
		// Check name collision // 检查名称冲突
		must.Null(res.mapName2Enum[enum.Name()])
		res.mapName2Enum[enum.Name()] = enum
		// Check basic collision // 检查 basic 枚举冲突
		must.Null(res.mapBasicEnum[enum.Basic()])
		res.mapBasicEnum[enum.Basic()] = enum
	}
	return res
}

// LookupByProto finds an Enum using its Protocol Buffer enum value
// Returns the Enum and true if found, nil and false otherwise
// Use this when you need to check existence before accessing the value
//
// 通过 Protocol Buffer 枚举值查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
// 当需要在访问值之前检查是否存在时使用此方法
func (c *Enums[P, B, M]) LookupByProto(proto P) (*Enum[P, B, M], bool) {
	if res, ok := c.mapProtoEnum[proto]; ok {
		return must.Full(res), true
	}
	return nil, false
}

// GetByProto finds an Enum using its Protocol Buffer enum value
// Uses the enum's numeric code when searching in the collection
// Returns default value if the enum is not found in the collection
// Panics if no default value has been configured
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 使用枚举的数字代码在集合中查找
// 如果在集合中找不到枚举则返回默认值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetByProto(proto P) *Enum[P, B, M] {
	if res, ok := c.mapProtoEnum[proto]; ok {
		return must.Full(res)
	}
	return c.GetDefault()
}

// MustGetByProto finds an Enum using its Protocol Buffer enum value
// Panics if the enum is not found in the collection
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 如果在集合中找不到枚举则会 panic
func (c *Enums[P, B, M]) MustGetByProto(proto P) *Enum[P, B, M] {
	return must.Nice(c.mapProtoEnum[proto])
}

// LookupByCode finds an Enum using its numeric code
// Returns the Enum and true if found, nil and false otherwise
// Use this when you need to check existence before accessing the value
//
// 通过数字代码查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
// 当需要在访问值之前检查是否存在时使用此方法
func (c *Enums[P, B, M]) LookupByCode(code int32) (*Enum[P, B, M], bool) {
	if res, ok := c.mapCode2Enum[code]; ok {
		return must.Full(res), true
	}
	return nil, false
}

// GetByCode finds an Enum using its numeric code
// Performs direct map lookup using the int32 code value
// Returns default value if no enum with the given code exists
// Panics if no default value has been configured
//
// 通过数字代码检索 Enum
// 使用 int32 代码值执行直接映射查找
// 如果不存在具有给定代码的枚举则返回默认值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetByCode(code int32) *Enum[P, B, M] {
	if res, ok := c.mapCode2Enum[code]; ok {
		return must.Full(res)
	}
	return c.GetDefault()
}

// MustGetByCode finds an Enum using its numeric code
// Panics if no enum with the given code exists
//
// 通过数字代码检索 Enum
// 如果不存在具有给定代码的枚举则会 panic
func (c *Enums[P, B, M]) MustGetByCode(code int32) *Enum[P, B, M] {
	return must.Nice(c.mapCode2Enum[code])
}

// LookupByName finds an Enum using its string name
// Returns the Enum and true if found, nil and false otherwise
// Use this when you need to check existence before accessing the value
//
// 通过字符串名称查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
// 当需要在访问值之前检查是否存在时使用此方法
func (c *Enums[P, B, M]) LookupByName(name string) (*Enum[P, B, M], bool) {
	if res, ok := c.mapName2Enum[name]; ok {
		return must.Full(res), true
	}
	return nil, false
}

// GetByName finds an Enum using its string name
// Performs direct map lookup using the enum name string
// Returns default value if no enum with the given name exists
// Panics if no default value has been configured
//
// 通过字符串名称检索 Enum
// 使用枚举名称字符串执行直接映射查找
// 如果不存在具有给定名称的枚举则返回默认值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetByName(name string) *Enum[P, B, M] {
	if res, ok := c.mapName2Enum[name]; ok {
		return must.Full(res)
	}
	return c.GetDefault()
}

// MustGetByName finds an Enum using its string name
// Panics if no enum with the given name exists
//
// 通过字符串名称检索 Enum
// 如果不存在具有给定名称的枚举则会 panic
func (c *Enums[P, B, M]) MustGetByName(name string) *Enum[P, B, M] {
	return must.Nice(c.mapName2Enum[name])
}

// LookupByBasic finds an Enum using its Go native enum value
// Returns the Enum and true if found, nil and false otherwise
// Use this when you need to check existence before accessing the value
//
// 通过 Go 原生枚举值查找 Enum
// 找到时返回 Enum 和 true，否则返回 nil 和 false
// 当需要在访问值之前检查是否存在时使用此方法
func (c *Enums[P, B, M]) LookupByBasic(basic B) (*Enum[P, B, M], bool) {
	if res, ok := c.mapBasicEnum[basic]; ok {
		return must.Full(res), true
	}
	return nil, false
}

// GetByBasic finds an Enum using its Go native enum value
// Performs direct map lookup using the basic enum value
// Returns default value if no enum with the given basic enum exists
// Panics if no default value has been configured
//
// 通过 Go 原生枚举值检索 Enum
// 使用 basic 枚举值执行直接映射查找
// 如果不存在具有给定 basic 枚举的枚举则返回默认值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetByBasic(basic B) *Enum[P, B, M] {
	if res, ok := c.mapBasicEnum[basic]; ok {
		return must.Full(res)
	}
	return c.GetDefault()
}

// MustGetByBasic finds an Enum using its Go native enum value
// Panics if no enum with the given basic enum exists
//
// 通过 Go 原生枚举值检索 Enum
// 如果不存在具有给定 basic 枚举的枚举则会 panic
func (c *Enums[P, B, M]) MustGetByBasic(basic B) *Enum[P, B, M] {
	return must.Nice(c.mapBasicEnum[basic])
}

// GetDefault returns the current default Enum value
// Panics if no default value has been configured
//
// 返回当前的默认 Enum 值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetDefault() *Enum[P, B, M] {
	return must.Full(c.defaultValue)
}

// ListProtos returns a slice containing each protoEnum value in the defined sequence
// Maintains the same sequence as enum values were registered
//
// 返回一个包含各 protoEnum 值的切片，次序与定义时一致
// 保持枚举值注册时的顺序
func (c *Enums[P, B, M]) ListProtos() []P {
	var results = make([]P, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		results = append(results, item.Proto())
	}
	return results
}

// ListBasics returns a slice containing each basicEnum value in the defined sequence
// Maintains the same sequence as enum values were registered
//
// 返回一个包含各 basicEnum 值的切片，次序与定义时一致
// 保持枚举值注册时的顺序
func (c *Enums[P, B, M]) ListBasics() []B {
	var results = make([]B, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		results = append(results, item.Basic())
	}
	return results
}

// ListValidProtos returns a slice excluding the default protoEnum value.
// If no default value is configured, returns each protoEnum value.
//
// 返回一个切片，排除默认 protoEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 protoEnum 值。
func (c *Enums[P, B, M]) ListValidProtos() []P {
	if c.defaultValue != nil && !utils.GetPointerValue(c.defaultValid) {
		var results []P
		for _, item := range c.enumElements {
			if item.Code() != c.defaultValue.Code() {
				results = append(results, item.Proto())
			}
		}
		return results
	}
	return c.ListProtos()
}

// ListValidBasics returns a slice excluding the default basicEnum value.
// If no default value is configured, returns each basicEnum value.
//
// 返回一个切片，排除默认 basicEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 basicEnum 值。
func (c *Enums[P, B, M]) ListValidBasics() []B {
	if c.defaultValue != nil && !utils.GetPointerValue(c.defaultValid) {
		var results []B
		for _, item := range c.enumElements {
			if item.Basic() != c.defaultValue.Basic() {
				results = append(results, item.Basic())
			}
		}
		return results
	}
	return c.ListBasics()
}
