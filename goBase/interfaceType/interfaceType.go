package interfacetype

import "fmt"

/*
interface(接口): 是 Go 语言中的一种类型，它定义了一组方法签名的集合。接口指定了要做什么，但不指定如何做。
接口的实现是隐式的，只要一个结构体实现了接口的所有方法，那么结构体就实现了该接口，接口也可以组合
*/

// 声明接口
type ReadInterface interface {
	read() string // 接口中方法之需要指定名称、传参数和返回字段信息和类型
}

type WriteInterface interface {
	write(s string)
}

// 声明结构体1
type ReadStruct struct {
}

// 实现接口
func (rs ReadStruct) read() string {
	return "hello!"
}

// 声明结构体2
type ReadStruct2 struct {
}

// 实现接口
func (rs ReadStruct2) read() string {
	return "hello!"
}

// 上述两个结构体都实现了接口那么就可以用接口名调用具体的方法
func InterfaceFun(inter ReadInterface) {
	// 接口调用Read1和Read2的read()方法
	fmt.Println("接口调用read1的结果是", inter.read())
}
