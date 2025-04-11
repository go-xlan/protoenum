package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protoenumstatus"
	"github.com/stretchr/testify/require"
)

var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)

func TestEnums_GetByEnum(t *testing.T) {
	enum := enums.GetByEnum(protoenumstatus.StatusEnum_SUCCESS)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Hans(), "成功")
}

func TestEnums_GetByCode(t *testing.T) {
	enum := enums.GetByCode(int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_FAILURE.String())
	require.Equal(t, enum.Hans(), "失败")
}

func TestEnums_GetByName(t *testing.T) {
	enum := enums.GetByName(protoenumstatus.StatusEnum_UNKNOWN.String())
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_UNKNOWN.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_UNKNOWN.String())
	require.Equal(t, enum.Hans(), "未知")
}

func TestEnums_GetByDesc(t *testing.T) {
	enum := enums.GetByDesc("成功")
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Hans(), "成功")
}

func TestEnums_GetByHans(t *testing.T) {
	enum := enums.GetByHans("成功")
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Hans(), "成功")
}
