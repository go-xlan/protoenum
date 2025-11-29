// Package protoenum: Collection management to handle Protocol Buffer enum metadata
// Provides indexed collections of enum descriptors with multiple lookup methods
// Enables fast lookup using code, name, and pure value with efficient enum handling
//
// protoenum: Protocol Buffer 枚举元数据集合管理
// 提供带有多种查找方法的枚举描述符索引集合
// 支持按代码、名称或朴素枚举值快速检索，实现高效枚举处理
package protoenum

import (
	"github.com/yyle88/must"
)

// Enums manages a collection of Enum instances with indexed lookups
// Maintains three maps enabling efficient lookup using different identifiers
// Provides O(1) lookup performance when searching code, name, and pure value
// Includes a configurable default value returned when lookups miss
//
// Enums 管理 Enum 实例集合并提供索引查找
// 维护三个映射表以通过不同标识符高效检索
// 为代码、名称和朴素枚举值搜索提供 O(1) 查找性能
// 支持在查找失败时返回可选的默认值
type Enums[protoEnum ProtoEnum, plainEnum comparable, extraMeta any] struct {
	mapCode2Enum map[int32]*Enum[protoEnum, plainEnum, extraMeta]     // Map from numeric code to Enum // 从数字代码到 Enum 的映射
	mapName2Enum map[string]*Enum[protoEnum, plainEnum, extraMeta]    // Map from name string to Enum // 从名称字符串到 Enum 的映射
	mapPure2Enum map[plainEnum]*Enum[protoEnum, plainEnum, extraMeta] // Map from plain enum to Enum // 从朴素枚举到 Enum 的映射
	defaultValue *Enum[protoEnum, plainEnum, extraMeta]               // Configurable default value when lookup misses // 查找失败时的可选默认值
}

// NewEnums creates a new Enums collection from the given Enum instances
// Builds indexed maps enabling efficient lookup using code, name, and pure value
// The first item becomes the default value if provided
// Returns a reference to the created Enums collection available when querying
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 构建索引映射以通过代码、名称和朴素枚举值高效查找
// 如果提供了参数，第一个项成为默认值
// 返回创建的 Enums 集合指针，准备好进行查询
func NewEnums[protoEnum ProtoEnum, plainEnum comparable, extraMeta any](params ...*Enum[protoEnum, plainEnum, extraMeta]) *Enums[protoEnum, plainEnum, extraMeta] {
	res := &Enums[protoEnum, plainEnum, extraMeta]{
		mapCode2Enum: make(map[int32]*Enum[protoEnum, plainEnum, extraMeta], len(params)),
		mapName2Enum: make(map[string]*Enum[protoEnum, plainEnum, extraMeta], len(params)),
		mapPure2Enum: make(map[plainEnum]*Enum[protoEnum, plainEnum, extraMeta], len(params)),
	}
	for _, enum := range params {
		must.Full(enum)

		// Check code collision // 检查代码冲突
		must.Null(res.mapCode2Enum[enum.Code()])
		res.mapCode2Enum[enum.Code()] = enum
		// Check name collision // 检查名称冲突
		must.Null(res.mapName2Enum[enum.Name()])
		res.mapName2Enum[enum.Name()] = enum
		// Check pure collision // 检查朴素枚举冲突
		must.Null(res.mapPure2Enum[enum.Pure()])
		res.mapPure2Enum[enum.Pure()] = enum
	}
	if len(params) > 0 {
		res.defaultValue = params[0] // Set first item as default if available // 如果有参数，将第一个设置为默认值
	}
	return res
}

// GetByEnum finds an Enum using its Protocol Buffer enum value
// Uses the enum's numeric code when searching in the collection
// Returns default value if the enum is not found in the collection
// Panics if no default value has been configured
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 使用枚举的数字代码在集合中查找
// 如果在集合中找不到枚举则返回默认值
// 如果未配置默认值则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) GetByEnum(enum protoEnum) *Enum[protoEnum, plainEnum, extraMeta] {
	if res, ok := c.mapCode2Enum[int32(enum.Number())]; ok {
		return must.Full(res)
	}
	return c.GetDefault()
}

// MustGetByEnum finds an Enum using its Protocol Buffer enum value
// Panics if the enum is not found in the collection
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 如果在集合中找不到枚举则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) MustGetByEnum(enum protoEnum) *Enum[protoEnum, plainEnum, extraMeta] {
	return must.Nice(c.mapCode2Enum[int32(enum.Number())])
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
func (c *Enums[protoEnum, plainEnum, extraMeta]) GetByCode(code int32) *Enum[protoEnum, plainEnum, extraMeta] {
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
func (c *Enums[protoEnum, plainEnum, extraMeta]) MustGetByCode(code int32) *Enum[protoEnum, plainEnum, extraMeta] {
	return must.Nice(c.mapCode2Enum[code])
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
func (c *Enums[protoEnum, plainEnum, extraMeta]) GetByName(name string) *Enum[protoEnum, plainEnum, extraMeta] {
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
func (c *Enums[protoEnum, plainEnum, extraMeta]) MustGetByName(name string) *Enum[protoEnum, plainEnum, extraMeta] {
	return must.Nice(c.mapName2Enum[name])
}

// GetByPure finds an Enum using its plain enum value
// Performs direct map lookup using the plain enum value
// Returns default value if no enum with the given plain enum exists
// Panics if no default value has been configured
//
// 通过朴素枚举值检索 Enum
// 使用朴素枚举值执行直接映射查找
// 如果不存在具有给定朴素枚举的枚举则返回默认值
// 如果未配置默认值则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) GetByPure(pure plainEnum) *Enum[protoEnum, plainEnum, extraMeta] {
	if res, ok := c.mapPure2Enum[pure]; ok {
		return must.Full(res)
	}
	return c.GetDefault()
}

// MustGetByPure finds an Enum using its plain enum value
// Panics if no enum with the given plain enum exists
//
// 通过朴素枚举值检索 Enum
// 如果不存在具有给定朴素枚举的枚举则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) MustGetByPure(pure plainEnum) *Enum[protoEnum, plainEnum, extraMeta] {
	return must.Nice(c.mapPure2Enum[pure])
}

// GetDefault returns the current default Enum value
// Panics if no default value has been configured
//
// 返回当前的默认 Enum 值
// 如果未配置默认值则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) GetDefault() *Enum[protoEnum, plainEnum, extraMeta] {
	return must.Full(c.defaultValue)
}

// SetDefault sets the default Enum value to return when lookups miss
// Allows dynamic configuration of the fallback value post creation
// Panics if defaultEnum is nil, use UnsetDefault to remove the default value
//
// 设置查找失败时返回的默认 Enum 值
// 允许在创建后动态配置回退值
// 如果 defaultEnum 为 nil 则会 panic，使用 UnsetDefault 清除默认值
func (c *Enums[protoEnum, plainEnum, extraMeta]) SetDefault(defaultEnum *Enum[protoEnum, plainEnum, extraMeta]) {
	must.Null(c.defaultValue)
	c.defaultValue = must.Full(defaultEnum)
}

// UnsetDefault unsets the default Enum value
// Once invoked, GetByXxx lookups panic if not found
// Panics if no default value exists at the moment
//
// 取消设置默认 Enum 值
// 调用此方法后，GetByXxx 查找失败时会 panic
// 如果当前无默认值则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) UnsetDefault() {
	must.Full(c.defaultValue)
	c.defaultValue = nil
}

// WithDefaultEnum sets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration during initialization
// Convenient when setting defaults in package-scope variable declarations
//
// 设置默认 Enum 值并返回 Enums 实例
// 支持初始化时的流式链式配置
// 适用于在全局变量声明中设置默认值
func (c *Enums[protoEnum, plainEnum, extraMeta]) WithDefaultEnum(defaultEnum *Enum[protoEnum, plainEnum, extraMeta]) *Enums[protoEnum, plainEnum, extraMeta] {
	c.SetDefault(defaultEnum)
	return c
}

// WithDefaultCode sets the default using a numeric code and returns the Enums instance
// Convenient chain method when you know the default enum code
// Panics if the specified code is not found in the collection
//
// 使用数字代码设置默认值并返回 Enums 实例
// 当你知道默认枚举代码时的便捷链式方法
// 如果指定的代码不存在则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) WithDefaultCode(code int32) *Enums[protoEnum, plainEnum, extraMeta] {
	return c.WithDefaultEnum(must.Full(c.mapCode2Enum[code]))
}

// WithDefaultName sets the default using an enum name and returns the Enums instance
// Convenient chain method when you know the default enum name
// Panics if the specified name is not found in the collection
//
// 使用枚举名称设置默认值并返回 Enums 实例
// 当你知道默认枚举名称时的便捷链式方法
// 如果指定的名称不存在则会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) WithDefaultName(name string) *Enums[protoEnum, plainEnum, extraMeta] {
	return c.WithDefaultEnum(must.Full(c.mapName2Enum[name]))
}

// WithUnsetDefault unsets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration to remove default value
// Once invoked, GetByXxx lookups panic if not found
//
// 取消设置默认 Enum 值并返回 Enums 实例
// 支持流式链式配置以移除默认值
// 之后 GetByXxx 查找失败时会 panic
func (c *Enums[protoEnum, plainEnum, extraMeta]) WithUnsetDefault() *Enums[protoEnum, plainEnum, extraMeta] {
	c.UnsetDefault()
	return c
}
