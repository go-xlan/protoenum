package protoenum

type Enums[protoEnum ProtoEnum] struct {
	mapCode2Enum map[int32]*Enum[protoEnum]
	mapName2Enum map[string]*Enum[protoEnum]
	mapDesc2Enum map[string]*Enum[protoEnum]
}

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
	return res
}

func (c *Enums[protoEnum]) GetByEnum(enum protoEnum) *Enum[protoEnum] {
	return c.mapCode2Enum[int32(enum.Number())]
}

func (c *Enums[protoEnum]) GetByCode(code int32) *Enum[protoEnum] {
	return c.mapCode2Enum[code]
}

func (c *Enums[protoEnum]) GetByName(name string) *Enum[protoEnum] {
	return c.mapName2Enum[name]
}

func (c *Enums[protoEnum]) GetByDesc(desc string) *Enum[protoEnum] {
	return c.mapDesc2Enum[desc]
}

func (c *Enums[protoEnum]) GetByHans(desc string) *Enum[protoEnum] {
	return c.mapDesc2Enum[desc]
}
