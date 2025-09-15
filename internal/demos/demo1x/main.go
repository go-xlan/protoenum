package main

import (
	"fmt"

	"github.com/go-xlan/protoenum"
	"github.com/go-xlan/protoenum/protos/protoenumstatus"
)

// Build status enum collection
var enums = protoenum.NewEnums(
	protoenum.NewEnum(protoenumstatus.StatusEnum_UNKNOWN, "未知"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_SUCCESS, "成功"),
	protoenum.NewEnum(protoenumstatus.StatusEnum_FAILURE, "失败"),
)

func main() {
	// Get enhanced description from protobuf enum (returns default if not found)
	successStatus := enums.GetByCode(int32(protoenumstatus.StatusEnum_SUCCESS))
	fmt.Printf("Status: %s\n", successStatus.Desc())

	// Convert between protoenum and native enum (safe with default fallback)
	statusEnum := enums.GetByName("SUCCESS")
	native := protoenumstatus.StatusEnum(statusEnum.Code())
	fmt.Printf("Native enum: %v\n", native)

	// Use in business logic
	if native == protoenumstatus.StatusEnum_SUCCESS {
		fmt.Println("Operation completed!")
	}
}
