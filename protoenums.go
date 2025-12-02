// Package protoenum: Collection management to handle Protocol Buffer enum metadata
// Provides indexed collections of enum descriptors with multiple lookup methods
// Enables fast lookup using code, name, and pure value with efficient enum handling
//
// protoenum: Protocol Buffer 枚举元数据集合管理
// 提供带有多种查找方法的枚举描述符索引集合
// 支持按代码、名称或朴素枚举值快速检索，实现高效枚举处理
package protoenum

import (
	"slices"

	"github.com/yyle88/must"
	"github.com/yyle88/tern/slicetern"
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
type Enums[P ProtoEnum, E comparable, M any] struct {
	enumElements []*Enum[P, E, M]          // Holds complete Enum instances in defined sequence // 存放所有 Enum 实例，并维持其定义的次序
	mapCode2Enum map[int32]*Enum[P, E, M]  // Map from numeric code to Enum // 从数字代码到 Enum 的映射
	mapName2Enum map[string]*Enum[P, E, M] // Map from name string to Enum // 从名称字符串到 Enum 的映射
	mapPure2Enum map[E]*Enum[P, E, M]      // Map from plain enum to Enum // 从朴素枚举到 Enum 的映射
	defaultValue *Enum[P, E, M]            // Configurable default value when lookup misses // 查找失败时的可选默认值
}

// NewEnums creates a new Enums collection from the given Enum instances
// Builds indexed maps enabling efficient lookup using code, name, and pure value
// The first item becomes the default value if provided
// Returns a reference to the created Enums collection, usable in lookup operations
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 构建索引映射以通过代码、名称和朴素枚举值高效查找
// 如果提供了参数，第一个项成为默认值
// 返回创建的 Enums 集合指针，可用于各种查找操作
func NewEnums[P ProtoEnum, E comparable, M any](params ...*Enum[P, E, M]) *Enums[P, E, M] {
	res := &Enums[P, E, M]{
		enumElements: slices.Clone(params), // Clone the slice to preserve the defined sequence of enum elements // 克隆切片以保持枚举元素的定义次序
		mapCode2Enum: make(map[int32]*Enum[P, E, M], len(params)),
		mapName2Enum: make(map[string]*Enum[P, E, M], len(params)),
		mapPure2Enum: make(map[E]*Enum[P, E, M], len(params)),
		defaultValue: slicetern.V0(params), // Set first item as default if available // 如果有参数，将第一个设置为默认值
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
func (c *Enums[P, E, M]) GetByEnum(enum P) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) MustGetByEnum(enum P) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) GetByCode(code int32) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) MustGetByCode(code int32) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) GetByName(name string) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) MustGetByName(name string) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) GetByPure(pure E) *Enum[P, E, M] {
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
func (c *Enums[P, E, M]) MustGetByPure(pure E) *Enum[P, E, M] {
	return must.Nice(c.mapPure2Enum[pure])
}

// GetDefault returns the current default Enum value
// Panics if no default value has been configured
//
// 返回当前的默认 Enum 值
// 如果未配置默认值则会 panic
func (c *Enums[P, E, M]) GetDefault() *Enum[P, E, M] {
	return must.Full(c.defaultValue)
}

// GetDefaultEnum returns the protoEnum value of the default Enum
// Panics if no default value has been configured
//
// 返回默认 Enum 的 protoEnum 值
// 如果未配置默认值则会 panic
func (c *Enums[P, E, M]) GetDefaultEnum() P {
	return must.Full(c.GetDefault()).Base()
}

// GetDefaultPure returns the plainEnum value of the default Enum
// Panics if no default value has been configured
//
// 返回默认 Enum 的 plainEnum 值
// 如果未配置默认值则会 panic
func (c *Enums[P, E, M]) GetDefaultPure() E {
	return must.Full(c.GetDefault()).Pure()
}

// SetDefault sets the default Enum value to return when lookups miss
// Allows dynamic configuration of the fallback value post creation
// Panics if defaultEnum is nil, use UnsetDefault to remove the default value
//
// 设置查找失败时返回的默认 Enum 值
// 允许在创建后动态配置回退值
// 如果 defaultEnum 为 nil 则会 panic，使用 UnsetDefault 清除默认值
func (c *Enums[P, E, M]) SetDefault(defaultEnum *Enum[P, E, M]) {
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
func (c *Enums[P, E, M]) UnsetDefault() {
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
func (c *Enums[P, E, M]) WithDefaultEnum(defaultEnum *Enum[P, E, M]) *Enums[P, E, M] {
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
func (c *Enums[P, E, M]) WithDefaultCode(code int32) *Enums[P, E, M] {
	return c.WithDefaultEnum(must.Full(c.mapCode2Enum[code]))
}

// WithDefaultName sets the default using an enum name and returns the Enums instance
// Convenient chain method when you know the default enum name
// Panics if the specified name is not found in the collection
//
// 使用枚举名称设置默认值并返回 Enums 实例
// 当你知道默认枚举名称时的便捷链式方法
// 如果指定的名称不存在则会 panic
func (c *Enums[P, E, M]) WithDefaultName(name string) *Enums[P, E, M] {
	return c.WithDefaultEnum(must.Full(c.mapName2Enum[name]))
}

// WithUnsetDefault unsets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration to remove default value
// Once invoked, GetByXxx lookups panic if not found
//
// 取消设置默认 Enum 值并返回 Enums 实例
// 支持流式链式配置以移除默认值
// 之后 GetByXxx 查找失败时会 panic
func (c *Enums[P, E, M]) WithUnsetDefault() *Enums[P, E, M] {
	c.UnsetDefault()
	return c
}

// ListEnums returns a slice containing each protoEnum value in the defined sequence.
// ListEnums 返回一个包含各 protoEnum 值的切片，次序与定义时一致。
func (c *Enums[P, E, M]) ListEnums() []P {
	var enums = make([]P, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		enums = append(enums, item.Base())
	}
	return enums
}

// ListPures returns a slice containing each plainEnum value in the defined sequence.
// ListPures 返回一个包含各 plainEnum 值的切片，次序与定义时一致。
func (c *Enums[P, E, M]) ListPures() []E {
	var pures = make([]E, 0, len(c.enumElements))
	for _, item := range c.enumElements {
		pures = append(pures, item.Pure())
	}
	return pures
}

// ListValidEnums returns a slice excluding the default protoEnum value.
// If no default value is configured, returns each protoEnum value.
//
// 返回一个切片，排除默认 protoEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 protoEnum 值。
func (c *Enums[P, E, M]) ListValidEnums() []P {
	if c.defaultValue != nil {
		var enums []P
		for _, item := range c.enumElements {
			if item != c.defaultValue {
				enums = append(enums, item.Base())
			}
		}
		return enums
	}
	return c.ListEnums()
}

// ListValidPures returns a slice excluding the default plainEnum value.
// If no default value is configured, returns each plainEnum value.
//
// 返回一个切片，排除默认 plainEnum 值，其余按定义次序排列。
// 如果未配置默认值，则返回所有 plainEnum 值。
func (c *Enums[P, E, M]) ListValidPures() []E {
	if c.defaultValue != nil {
		var pures []E
		for _, item := range c.enumElements {
			if item != c.defaultValue {
				pures = append(pures, item.Pure())
			}
		}
		return pures
	}
	return c.ListPures()
}
