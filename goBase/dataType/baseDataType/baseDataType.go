package basedatatype

import "fmt"

/*
go提供的基本数据类型包括unit, unit8...unit64, int, int8...int64(int相关类型), rune(中文字符), byte(二进制), float32, float64, bool, string

	go中int8和int不是一个类型
	int类型(int， int8...int64) 占位符%d, 默认值是0，int == int64
	rune(中文字符，占3个字节) 占位符%c, 默认值是''
	byte(二进制,也可以表示英文状态下的一个字节) 占位符%b, 默认值是''
	string 占位符%s, 默认值是""
	bool 占位符%t, 默认值是 false
	float32/float64 占位符%f,可以在 f 前面添加数字控制小数位，默认值是0.000000 默认类型是float64
	%v 可以表示任意数据类型
	%T 获取数据类型
	%#v 以go数据形式展示,常用于表示空字符串
	%p 表示指针数据类型
*/
func BaseDataType() {
	var name string

	// invalid operation: age == num (mismatched types int and int8)
	var age int

	var num int8

	var score float64

	var gender bool

	var tip byte

	var title rune

	fmt.Printf("name: %s, age: %d, num: %d, score: %f, gender: %t, tip: %b, title: %c\n", name, age, num, score, gender, tip, title)
}
