package protoenum_test

import (
	"testing"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
	"github.com/stretchr/testify/require"
)

// TestEnums_GetByProto tests lookup with Protocol Buffer enum value
// Checks that GetByProto returns the correct Enum with matching properties
//
// 验证通过 Protocol Buffer 枚举值检索
// 测试 GetByProto 返回具有匹配属性的正确 Enum
func TestEnums_GetByProto(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	enum := enums.GetByProto(protoenumstatus.StatusEnum_SUCCESS)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Basic(), StatusTypeSuccess)
}

// TestEnums_GetByCode tests lookup using numeric code
// Checks that GetByCode returns the correct Enum using int32 code
//
// 验证通过数字代码检索
// 测试 GetByCode 使用 int32 代码返回正确的 Enum
func TestEnums_GetByCode(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	enum := enums.GetByCode(int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_FAILURE.String())
	require.Equal(t, enum.Basic(), StatusTypeFailure)
}

// TestEnums_GetByName tests lookup using string name
// Checks that GetByName returns the correct Enum using enum name
//
// 验证通过字符串名称检索
// 测试 GetByName 使用枚举名称返回正确的 Enum
func TestEnums_GetByName(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	enum := enums.GetByName(protoenumstatus.StatusEnum_UNKNOWN.String())
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_UNKNOWN.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_UNKNOWN.String())
	require.Equal(t, enum.Basic(), StatusTypeUnknown)
}

// TestEnums_GetByBasic verifies lookup using Go native enum value
// Tests that GetByBasic returns the correct Enum using basic enum type
//
// 验证通过 Go 原生枚举值检索
// 测试 GetByBasic 使用 basic 枚举类型返回正确的 Enum
func TestEnums_GetByBasic(t *testing.T) {
	type StatusType string
	const (
		StatusTypeUnknown StatusType = "unknown"
		StatusTypeSuccess StatusType = "success"
		StatusTypeFailure StatusType = "failure"
	)

	enums := protoenum.NewEnums(
		protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, StatusTypeUnknown),
		protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, StatusTypeSuccess),
		protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, StatusTypeFailure),
	)

	// Lookup using Go native enum value
	enum := enums.GetByBasic(StatusTypeSuccess)
	require.NotNil(t, enum)
	t.Log(enum.Code())
	t.Log(enum.Name())
	t.Log(enum.Basic())

	require.Equal(t, enum.Code(), int32(protoenumstatus.StatusEnum_SUCCESS.Number()))
	require.Equal(t, enum.Name(), protoenumstatus.StatusEnum_SUCCESS.String())
	require.Equal(t, enum.Basic(), StatusTypeSuccess)

	// Check MustGetByBasic works too
	enumFailure := enums.MustGetByBasic(StatusTypeFailure)
	require.Equal(t, enumFailure.Code(), int32(protoenumstatus.StatusEnum_FAILURE.Number()))
	require.Equal(t, enumFailure.Basic(), StatusTypeFailure)
}

// TestEnums_ListProtos tests the ListProtos method returns each protoEnum value
// Verifies that the returned slice maintains the defined sequence
//
// 测试 ListProtos 方法返回所有 protoEnum 值
// 验证返回的切片保持定义时的次序
func TestEnums_ListProtos(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	var enums = protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	protoEnums := enums.ListProtos()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 4)
	require.Equal(t, protoenumresult.ResultEnum_UNKNOWN, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[2])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[3])
}

// TestEnums_ListBasics tests the ListBasics method returns each basicEnum value
// Verifies that the returned slice maintains the defined sequence
//
// 测试 ListBasics 方法返回所有 basicEnum 值
// 验证返回的切片保持定义时的次序
func TestEnums_ListBasics(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	var enums = protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	basicEnums := enums.ListBasics()
	t.Log(basicEnums)
	require.Len(t, basicEnums, 4)
	require.Equal(t, ResultTypeUnknown, basicEnums[0])
	require.Equal(t, ResultTypePass, basicEnums[1])
	require.Equal(t, ResultTypeMiss, basicEnums[2])
	require.Equal(t, ResultTypeSkip, basicEnums[3])
}

// TestEnums_ListValidProtos tests the ListValidProtos method
// Verifies that the returned slice excludes the default protoEnum value
//
// 测试 ListValidProtos 方法
// 验证返回的切片排除默认 protoEnum 值
func TestEnums_ListValidProtos(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	var enums = protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	// Default is UNKNOWN (first item), so result should exclude it
	protoEnums := enums.ListValidProtos()
	t.Log(protoEnums)
	require.Len(t, protoEnums, 3)
	require.Equal(t, protoenumresult.ResultEnum_PASS, protoEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_MISS, protoEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, protoEnums[2])

	// When no default is set, returns each value
	enums.UnsetDefault()
	allEnums := enums.ListValidProtos()
	t.Log(allEnums)
	require.Len(t, allEnums, 4)
	require.Equal(t, protoenumresult.ResultEnum_UNKNOWN, allEnums[0])
	require.Equal(t, protoenumresult.ResultEnum_PASS, allEnums[1])
	require.Equal(t, protoenumresult.ResultEnum_MISS, allEnums[2])
	require.Equal(t, protoenumresult.ResultEnum_SKIP, allEnums[3])
}

// TestEnums_ListValidBasics tests the ListValidBasics method
// Verifies that the returned slice excludes the default basicEnum value
//
// 测试 ListValidBasics 方法
// 验证返回的切片排除默认 basicEnum 值
func TestEnums_ListValidBasics(t *testing.T) {
	type ResultType string

	const (
		ResultTypeUnknown ResultType = "unknown"
		ResultTypePass    ResultType = "pass"
		ResultTypeMiss    ResultType = "miss"
		ResultTypeSkip    ResultType = "skip"
	)

	var enums = protoenum.NewEnums(
		protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, ResultTypeUnknown),
		protoenum.NewEnum(protoenumresult.ResultEnum_PASS, ResultTypePass),
		protoenum.NewEnum(protoenumresult.ResultEnum_MISS, ResultTypeMiss),
		protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, ResultTypeSkip),
	)

	// Default is UNKNOWN (first item), so result should exclude it
	validBasics := enums.ListValidBasics()
	t.Log(validBasics)
	require.Len(t, validBasics, 3)
	require.Equal(t, ResultTypePass, validBasics[0])
	require.Equal(t, ResultTypeMiss, validBasics[1])
	require.Equal(t, ResultTypeSkip, validBasics[2])

	// When no default is set, returns each value
	enums.UnsetDefault()
	allBasics := enums.ListValidBasics()
	t.Log(allBasics)
	require.Len(t, allBasics, 4)
	require.Equal(t, ResultTypeUnknown, allBasics[0])
	require.Equal(t, ResultTypePass, allBasics[1])
	require.Equal(t, ResultTypeMiss, allBasics[2])
	require.Equal(t, ResultTypeSkip, allBasics[3])
}
