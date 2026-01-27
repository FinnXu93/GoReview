package scan

import (
	"fmt"
)

func ScanFun() {
	// Scan
	fmt.Println("请输入年龄:")

	var age int
	fmt.Scan(&age)

	fmt.Println("输入的年龄是", age)
}
