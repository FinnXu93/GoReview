package variable

import "fmt"

func VariableFun() {
	// 1.1 声明变量
	var n1 string
	// 1.2 赋值
	n1 = "finn"

	// 2 省略类型，声明加赋值
	var n2 = "chaha"

	// 3. 冒等式声明并赋值变量
	n3 := "chuchu"

	fmt.Printf("n1: %s, n2: %s, n3: %s\n", n1, n2, n3)

}
