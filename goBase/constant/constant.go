package constant

import "fmt"

// 常量声明时必须赋值，只读，不能使用冒等式赋值
const name string = "finn"

func GetConst() {
	fmt.Println("常量name的值是", name)
}
