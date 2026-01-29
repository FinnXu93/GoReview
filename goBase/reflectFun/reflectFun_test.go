package reflectfun

import (
	"testing"
)

type str1 struct {
	Name string `json:"name"`
	Age  uint8  `json:"age" title:"年龄"`
}

func TestReflectGetFun(t *testing.T) {

	// 数组
	var Arr1 [3]int8 = [3]int8{1, 2, 3}
	ReflectGetFun(Arr1)

	var map1 map[int]string = map[int]string{1: "a", 2: "b"}
	ReflectGetFun(map1)

	var s1 string = "abcdef"
	ReflectGetFun(s1)

	ReflectGetFun(str1{"finn", 13})

}
