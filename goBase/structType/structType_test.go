package structtype

import (
	"fmt"
	"testing"
)

func TestStructTypeFun(t *testing.T) {
	StructTypeFun()
}

func TestSetStructAddress(t *testing.T) {
	test1 := DetailedStruct{Address: "11111"}
	SetStructAddress(&test1, "33333")
	fmt.Println("new add:", test1.Address)

	fmt.Println("调用SimpleInfo中getAllInfo的结果是", test1.getAllInfo())
}
