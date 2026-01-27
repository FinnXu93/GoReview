package basedatatype

import (
	"fmt"
	"strconv"
)

// go不同python，类型转换不是隐式的，需要显示转换
func IntTransfun() {
	// int转其他
	var age int8 = 12
	// string()方法转换int时，会将整数换成对应的unicode码点
	// fmt.Println("int转string:", string(age)) 不能这么用
	// int转字符串
	// 方法1 基于fmt.Sprintf()方法
	stringAge := fmt.Sprintf("%d", age)
	fmt.Println("fmt.Sprintf()方法, int转string:", stringAge)

	// 方法2 基于strconv包的formatInt()
	stringAge2 := strconv.FormatInt(int64(age), 10)
	fmt.Println("strconv.formatInt()方法, int转string:", stringAge2)

	// int不能转bool：cannot convert age (variable of type int8) to type bool
	fmt.Println("int转float64:", float64(age))

	fmt.Println("int转byte:", byte(age))

	fmt.Println("int转rune:", rune(age))

}

func BoolTransFun() {
	var gender bool = true

	// bool不能和int互转
	// fmt.Println("bool 转 int:", int(gender))
	// fmt.Println("bool 转 int:", int64(gender))

	// bool 转字符串, 不能使用string(bool)
	strBool := fmt.Sprintf("%t", gender)
	fmt.Println("fmt.Sprintf(), bool 转 string:", strBool)

	strBool2 := strconv.FormatBool(gender)
	fmt.Println("fstrconv.formatBool(), bool 转 string:", strBool2)

}
