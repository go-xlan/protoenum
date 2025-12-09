package protoenum

import "github.com/yyle88/must"

// GetDefaultProto returns the protoEnum value of the default Enum
// Panics if no default value has been configured
//
// 返回默认 Enum 的 protoEnum 值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetDefaultProto() P {
	return must.Full(c.GetDefault()).Proto()
}

// GetDefaultBasic returns the basicEnum value of the default Enum
// Panics if no default value has been configured
//
// 返回默认 Enum 的 basicEnum 值
// 如果未配置默认值则会 panic
func (c *Enums[P, B, M]) GetDefaultBasic() B {
	return must.Full(c.GetDefault()).Basic()
}

// SetDefault sets the default Enum value to return when lookups miss
// Allows dynamic configuration of the fallback value post creation
// Panics if defaultEnum is nil, use UnsetDefault to remove the default value
//
// 设置查找失败时返回的默认 Enum 值
// 允许在创建后动态配置回退值
// 如果 defaultEnum 为 nil 则会 panic，使用 UnsetDefault 清除默认值
func (c *Enums[P, B, M]) SetDefault(enum *Enum[P, B, M]) {
	must.Null(c.defaultValue)
	// Note: use SetDefaultProto and SetDefaultBasic to validate against map
	// 注意：使用 SetDefaultProto 或 SetDefaultBasic 可确保值在集合中存在
	c.defaultValue = must.Full(enum)
	// Proto default is often treated as non-active to distinguish set from unset
	// Proto 默认值通常无效，以区分已设置和未设置的情况
	c.defaultValid = nil
}

// SetDefaultProto sets the default using a Protocol Buffer enum value
// Panics if the specified proto enum is not found in the collection
//
// 使用 Protocol Buffer 枚举值设置默认值
// 如果指定的 proto 枚举不存在则会 panic
func (c *Enums[P, B, M]) SetDefaultProto(proto P) {
	c.SetDefault(c.MustGetByProto(proto))
}

// SetDefaultBasic sets the default using a Go native enum value
// Panics if the specified basic enum is not found in the collection
//
// 使用 Go 原生枚举值设置默认值
// 如果指定的 basic 枚举不存在则会 panic
func (c *Enums[P, B, M]) SetDefaultBasic(basic B) {
	c.SetDefault(c.MustGetByBasic(basic))
}

// SetDefaultValid marks the default value as active when true
// When active, ListValidProtos and ListValidBasics include the default
// Panics if no default value exists, panics if defaultValid has been set
//
// 标记默认值是否应被视为有效
// 当 valid 为 true 时，ListValidProtos 和 ListValidBasics 包含默认值
// 如果无默认值或 defaultValid 已设置则会 panic
func (c *Enums[P, B, M]) SetDefaultValid(valid bool) {
	must.Full(c.defaultValue)
	must.Null(c.defaultValid)
	c.defaultValid = &valid
}

// UnsetDefault unsets the default Enum value
// Once invoked, GetByXxx lookups panic if not found
// Panics if no default value exists at the moment
//
// 取消设置默认 Enum 值
// 调用此方法后，GetByXxx 查找失败时会 panic
// 如果当前无默认值则会 panic
func (c *Enums[P, B, M]) UnsetDefault() {
	must.Full(c.defaultValue)
	c.defaultValue = nil
	c.defaultValid = nil
}

// WithDefault sets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration during initialization
// Convenient when setting defaults in package-scope variable declarations
//
// 设置默认 Enum 值并返回 Enums 实例
// 支持初始化时的流式链式配置
// 适用于在全局变量声明中设置默认值
func (c *Enums[P, B, M]) WithDefault(enum *Enum[P, B, M]) *Enums[P, B, M] {
	c.SetDefault(enum)
	return c
}

// WithDefaultProto sets the default using a proto enum and returns the Enums instance
// Convenient chain method when you know the default Protocol Buffer enum value
// Panics if the specified proto enum is not found in the collection
//
// 使用 proto 枚举设置默认值并返回 Enums 实例
// 当你知道默认 Protocol Buffer 枚举值时的便捷链式方法
// 如果指定的 proto 枚举不存在则会 panic
func (c *Enums[P, B, M]) WithDefaultProto(proto P) *Enums[P, B, M] {
	c.SetDefaultProto(proto)
	return c
}

// WithDefaultBasic sets the default using a basic enum and returns the Enums instance
// Convenient chain method when you know the default Go native enum value
// Panics if the specified basic enum is not found in the collection
//
// 使用 basic 枚举设置默认值并返回 Enums 实例
// 当你知道默认 Go 原生枚举值时的便捷链式方法
// 如果指定的 basic 枚举不存在则会 panic
func (c *Enums[P, B, M]) WithDefaultBasic(basic B) *Enums[P, B, M] {
	c.SetDefaultBasic(basic)
	return c
}

// WithDefaultCode sets the default using a numeric code and returns the Enums instance
// Convenient chain method when you know the default enum code
// Panics if the specified code is not found in the collection
//
// 使用数字代码设置默认值并返回 Enums 实例
// 当你知道默认枚举代码时的便捷链式方法
// 如果指定的代码不存在则会 panic
func (c *Enums[P, B, M]) WithDefaultCode(code int32) *Enums[P, B, M] {
	return c.WithDefault(must.Full(c.mapCode2Enum[code]))
}

// WithDefaultName sets the default using an enum name and returns the Enums instance
// Convenient chain method when you know the default enum name
// Panics if the specified name is not found in the collection
//
// 使用枚举名称设置默认值并返回 Enums 实例
// 当你知道默认枚举名称时的便捷链式方法
// 如果指定的名称不存在则会 panic
func (c *Enums[P, B, M]) WithDefaultName(name string) *Enums[P, B, M] {
	return c.WithDefault(must.Full(c.mapName2Enum[name]))
}

// WithDefaultValid marks the default as active and returns the Enums instance
// When active, ListValidProtos and ListValidBasics include the default
// Convenient chain method to configure default active state at initialization
//
// 标记默认值为有效并返回 Enums 实例
// 当 valid 为 true 时，ListValidProtos 和 ListValidBasics 包含默认值
// 用于初始化时配置默认值有效性的便捷链式方法
func (c *Enums[P, B, M]) WithDefaultValid(valid bool) *Enums[P, B, M] {
	c.SetDefaultValid(valid)
	return c
}

// WithUnsetDefault unsets the default Enum value and returns the Enums instance
// Enables fluent chain-style configuration to remove default value
// Once invoked, GetByXxx lookups panic if not found
//
// 取消设置默认 Enum 值并返回 Enums 实例
// 支持流式链式配置以移除默认值
// 之后 GetByXxx 查找失败时会 panic
func (c *Enums[P, B, M]) WithUnsetDefault() *Enums[P, B, M] {
	c.UnsetDefault()
	return c
}
