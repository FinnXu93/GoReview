package interfacetype

import "testing"

func TestInterfaceFun(t *testing.T) {
	var inters []ReadInterface = []ReadInterface{
		ReadStruct{},
		ReadStruct2{},
	}
	for _, str := range inters {
		InterfaceFun(str)

	}
}
