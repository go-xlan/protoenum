package protoenum

// MetaNone represents blank metadata when enums have no description
//
// MetaNone 代表无描述枚举的空元数据
type MetaNone struct{}

// MetaDesc represents metadata with string description attached to enums
//
// MetaDesc 代表带字符串描述的枚举元数据
type MetaDesc struct{ description string }

// Desc returns the custom description of the enum
// Provides human-readable description with documentation purposes
//
// 返回枚举的自定义描述
// 提供人类可读的描述用于文档目的
func (c *MetaDesc) Desc() string {
	return c.description
}
