package main

import (
	"fmt"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumresult"
)

// Build enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumresult.ResultEnum_UNKNOWN, "其它"),
	protoenum.NewEnum(protoenumresult.ResultEnum_PASS, "通过"),
	protoenum.NewEnum(protoenumresult.ResultEnum_FAIL, "出错"),
	protoenum.NewEnum(protoenumresult.ResultEnum_SKIP, "跳过"),
)

func main() {
	// Lookup by enum code (returns default if not found)
	skipResult := enums.GetByCode(int32(protoenumresult.ResultEnum_SKIP))
	fmt.Printf("Result: %s\n", skipResult.Desc())

	// Lookup by enum name (safe with default fallback)
	passResult := enums.GetByName("PASS")
	native := protoenumresult.ResultEnum(passResult.Code())
	fmt.Printf("Native: %v\n", native)

	// Business logic with native enum
	if native == protoenumresult.ResultEnum_PASS {
		fmt.Println("Test passed!")
	}

	// Lookup by Chinese description (returns default if not found)
	result := enums.GetByDesc("跳过")
	fmt.Printf("Name: %s\n", result.Name())
}
