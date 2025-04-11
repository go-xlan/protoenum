package protoenum

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ProtoEnum interface {
	String() string
	Number() protoreflect.EnumNumber
}

type Enum[protoEnum ProtoEnum] struct {
	enum        protoEnum
	description string
}

func NewEnum[protoEnum ProtoEnum](enum protoEnum, description string) *Enum[protoEnum] {
	return &Enum[protoEnum]{
		enum:        enum,
		description: description,
	}
}

func (c *Enum[protoEnum]) Base() protoEnum {
	return c.enum
}

func (c *Enum[protoEnum]) Code() int32 {
	return int32(c.enum.Number())
}

func (c *Enum[protoEnum]) Name() string {
	return c.enum.String()
}

func (c *Enum[protoEnum]) Desc() string {
	return c.description
}

func (c *Enum[protoEnum]) Hans() string {
	return c.description
}
