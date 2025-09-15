// Package protoenum: Collection management handling Protocol Buffer enum metadata
// Provides indexed collections of enum descriptors with multiple lookup methods
// Enables fast lookup using code, name, and description with efficient enum handling
//
// protoenum: Protocol Buffer 枚举元数据集合管理
// 提供带有多种查找方法的枚举描述符索引集合
// 支持按代码、名称或描述快速检索，实现高效枚举处理
package protoenum

import (
	"github.com/pkg/errors"
)

// Enums manages a collection of Enum instances with indexed lookups
// Maintains three maps enabling efficient lookup using different identifiers
// Provides O(1) lookup performance when searching code, name, and description
// Supports optional default value when lookups return nothing
//
// Enums 管理 Enum 实例集合并提供索引查找
// 维护三个映射表以通过不同标识符高效检索
// 为代码、名称和描述搜索提供 O(1) 查找性能
// 支持在查找失败时返回可选的默认值
type Enums[protoEnum ProtoEnum] struct {
	mapCode2Enum map[int32]*Enum[protoEnum]  // Map from numeric code to Enum // 从数字代码到 Enum 的映射
	mapName2Enum map[string]*Enum[protoEnum] // Map from name string to Enum // 从名称字符串到 Enum 的映射
	mapDesc2Enum map[string]*Enum[protoEnum] // Map from description to Enum // 从描述到 Enum 的映射
	defaultValue *Enum[protoEnum]            // Optional default value when lookup fails // 查找失败时的可选默认值
}

// NewEnums creates a new Enums collection from the given Enum instances
// Builds indexed maps enabling efficient lookup using code, name, and description
// The first item is automatically set as the default value if provided
// Returns a reference to the created Enums collection available when querying
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 构建索引映射以通过代码、名称和描述高效查找
// 如果提供了参数，第一个项自动设置为默认值
// 返回创建的 Enums 集合指针，准备好进行查询
func NewEnums[protoEnum ProtoEnum](params ...*Enum[protoEnum]) *Enums[protoEnum] {
	res := &Enums[protoEnum]{
		mapCode2Enum: make(map[int32]*Enum[protoEnum], len(params)),
		mapName2Enum: make(map[string]*Enum[protoEnum], len(params)),
		mapDesc2Enum: make(map[string]*Enum[protoEnum], len(params)),
	}
	for _, enum := range params {
		res.mapCode2Enum[enum.Code()] = enum
		res.mapName2Enum[enum.Name()] = enum
		res.mapDesc2Enum[enum.Desc()] = enum
	}
	if len(params) > 0 {
		res.defaultValue = params[0] // Set first item as default if available // 如果有参数，将第一个设置为默认值
	}
	return res
}

// GetByEnum finds an Enum using its Protocol Buffer enum value
// Uses the enum's numeric code when searching in the collection
// Returns default value if the enum is not found in the collection
//
// 通过 Protocol Buffer 枚举值检索 Enum
// 使用枚举的数字代码在集合中查找
// 如果在集合中找不到枚举则返回默认值
func (c *Enums[protoEnum]) GetByEnum(enum protoEnum) *Enum[protoEnum] {
	if result := c.mapCode2Enum[int32(enum.Number())]; result != nil {
		return result
	}
	return c.defaultValue
}

// GetByCode finds an Enum using its numeric code
// Performs direct map lookup using the int32 code value
// Returns default value if no enum with the given code exists
//
// 通过数字代码检索 Enum
// 使用 int32 代码值执行直接映射查找
// 如果不存在具有给定代码的枚举则返回默认值
func (c *Enums[protoEnum]) GetByCode(code int32) *Enum[protoEnum] {
	if result := c.mapCode2Enum[code]; result != nil {
		return result
	}
	return c.defaultValue
}

// GetByName finds an Enum using its string name
// Performs direct map lookup using the enum name string
// Returns default value if no enum with the given name exists
//
// 通过字符串名称检索 Enum
// 使用枚举名称字符串执行直接映射查找
// 如果不存在具有给定名称的枚举则返回默认值
func (c *Enums[protoEnum]) GetByName(name string) *Enum[protoEnum] {
	if result := c.mapName2Enum[name]; result != nil {
		return result
	}
	return c.defaultValue
}

// GetByDesc finds an Enum using its description
// Performs direct map lookup using the custom description string
// Returns default value if no enum with the given description exists
//
// 通过描述检索 Enum
// 使用自定义描述字符串执行直接映射查找
// 如果不存在具有给定描述的枚举则返回默认值
func (c *Enums[protoEnum]) GetByDesc(desc string) *Enum[protoEnum] {
	if result := c.mapDesc2Enum[desc]; result != nil {
		return result
	}
	return c.defaultValue
}

// GetByHans finds an Enum using its Chinese description
// Alias to GetByDesc method, convenient with Chinese language support
// Returns default value if no enum with the given Chinese description exists
//
// 通过中文描述检索 Enum
// GetByDesc 方法的别名，方便中文语言支持
// 如果不存在具有给定中文描述的枚举则返回默认值
func (c *Enums[protoEnum]) GetByHans(desc string) *Enum[protoEnum] {
	return c.GetByDesc(desc)
}

// SetDefault sets the default Enum value to return when lookups fail
// Allows dynamic configuration of the fallback value after creation
// Pass nil to clear the default value
//
// 设置查找失败时返回的默认 Enum 值
// 允许在创建后动态配置回退值
// 传递 nil 可清除默认值
func (c *Enums[protoEnum]) SetDefault(defaultEnum *Enum[protoEnum]) {
	c.defaultValue = defaultEnum
}

// WithDefaultEnum sets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration during initialization
// Useful for setting defaults in global variable declarations
//
// 设置默认 Enum 值并返回 Enums 实例
// 支持初始化时的流式链式配置
// 适用于在全局变量声明中设置默认值
func (c *Enums[protoEnum]) WithDefaultEnum(defaultEnum *Enum[protoEnum]) *Enums[protoEnum] {
	c.defaultValue = defaultEnum
	return c
}

// WithDefaultCode sets the default using a numeric code and returns the Enums instance
// Convenient chain method when you know the default enum code
// Panics if the specified code is not found in the collection
//
// 使用数字代码设置默认值并返回 Enums 实例
// 当你知道默认枚举代码时的便捷链式方法
// 如果指定的代码不存在则会 panic
func (c *Enums[protoEnum]) WithDefaultCode(code int32) *Enums[protoEnum] {
	enum := c.mapCode2Enum[code]
	if enum == nil {
		panic(errors.Errorf("enum with code %d not found in collection", code))
	}
	c.defaultValue = enum
	return c
}

// WithDefaultName sets the default using an enum name and returns the Enums instance
// Convenient chain method when you know the default enum name
// Panics if the specified name is not found in the collection
//
// 使用枚举名称设置默认值并返回 Enums 实例
// 当你知道默认枚举名称时的便捷链式方法
// 如果指定的名称不存在则会 panic
func (c *Enums[protoEnum]) WithDefaultName(name string) *Enums[protoEnum] {
	enum := c.mapName2Enum[name]
	if enum == nil {
		panic(errors.Errorf("enum with name %s not found in collection", name))
	}
	c.defaultValue = enum
	return c
}

// GetDefault returns the current default Enum value
// Returns nil if no default value has been configured
//
// 返回当前的默认 Enum 值
// 如果未配置默认值则返回 nil
func (c *Enums[protoEnum]) GetDefault() *Enum[protoEnum] {
	return c.defaultValue
}
