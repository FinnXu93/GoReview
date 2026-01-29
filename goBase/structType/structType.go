package structtype

import "fmt"

// 声明结构体
type SiplemStruct1 struct {
	Name  string `json:"name" title:"姓名"` // ``提供结构体标签，用于json化、web字段校验等
	Age   uint8  `json:"age"`
	Phone int    `json:"omitempty"` // `omitempty表示json化时忽略`
}

func (ss SiplemStruct1) getAllInfo() map[string]any {
	var result map[string]any = make(map[string]any)
	result["name"] = ss.Name
	result["age"] = ss.Age
	return result
}

type SimpleStruct2 struct {
	Work string
}

type DetailedStruct struct {
	SiplemStruct1
	SimpleStruct2        // 结构体组合, 在当前结构体中能使用被组合的结构体的属性和方法，是has的关系
	Address       string `json:"address"`
}

// 结构体绑定方法
// 获取地址
func (de DetailedStruct) getAddress() string {
	return de.Address
}

// 修改结构中某个属性值
func (de *DetailedStruct) setAddress(newAddress string) {
	de.Address = newAddress
}

func StructTypeFun() {
	/*
				struct是 Go 语言中一种自定义的复合数据类型，
				它允许你将多个不同类型的数据字段组合成一个单一的、有意义的实体

				为什么要有结构体？
				1. 变量之间没有逻辑关联
				2. 难以传递一组相关数据
		 		3. 代码可读性差
				4. 容易出错（可能修改错误的变量）

	*/
	// 1.1 实例结构体
	var oneInfo DetailedStruct
	// 1.2 赋值
	oneInfo = DetailedStruct{SiplemStruct1{Name: "finn", Age: 12}, SimpleStruct2{Work: "it"}, "xxxxxx"}

	// 2 实例化结构体并赋值
	twoInfo := DetailedStruct{
		SiplemStruct1: SiplemStruct1{
			Name: "chaha",
			Age:  15,
		},
		Address: "xxxxxx",
	}

	// 结构体取值
	fmt.Println("oneInfo 中 Name 是", oneInfo.Name)

	// 修改结构体某个值
	twoInfo.Age = 13
	fmt.Println("修改twoInfo中Age后是", twoInfo)

	// 调用结构体方法
	fmt.Println("获取到的 address是", oneInfo.getAddress())

	oneInfo.setAddress("222222")

	fmt.Println("修改 oneInfo address 后是", oneInfo)

}

// 结构体作为函数参数传递是值传递
func SetStructAddress(st *DetailedStruct, newAdd string) {
	st.Address = newAdd
	fmt.Println("SetStructAddress 中 address 是", st.Address)
}
