package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protoenumstatus"
	"github.com/stretchr/testify/require"
)

func TestNewEnum(t *testing.T) {
	enum := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "任务完成")
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Desc())
	t.Log(enum.Hans())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Desc(), "任务完成")
	require.Equal(t, enum.Hans(), "任务完成")
}

func TestEnum_Base(t *testing.T) {
	enum := protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "审批通过")
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Desc())
	t.Log(enum.Hans())

	statusEnum := enum.Base()

	t.Log(statusEnum.String())
	t.Log(statusEnum.Number())
	require.Equal(t, statusEnum.String(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, statusEnum.Number(), protoenumstatus.StatusEnum_SUCCESS.Number())

	t.Log(statusEnum.Type().Descriptor().Name())
	require.Equal(t, statusEnum.Type().Descriptor().Name(), protoenumstatus.StatusEnum_SUCCESS.Type().Descriptor().Name())
}
