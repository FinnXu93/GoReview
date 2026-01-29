package complexdatatype

import (
	"fmt"
)

// 切片、映射(map)、指针、chan(通道)、函数、接口都是引用传递，剩下的都是值传递
/*
	修改后的slice1是 [2 3 4 100 1 100 12 100]
	修改后的slice1的指向数组的内存地址是0x1400011a200
	原始slice1是 [2 3 4 100 1 100 12 100]
	原始slice1的指向数组的内存地址是0x1400011a200
*/

// 数组
func ArrayFun() {
	/* 数组：一个固定长度且元素类型相同的数据集合, 类型表示: [n]T
	// 数组特点: 声明时必须告知长度，元素字能是一个类型， 元素默认值是0
	// 数组方法：
	// 	1. 赋值 var arr1 [3]int8 = [3]int8{1, 2, 3}
		2. 基于下标数组取值,注意索引越界 num := arr1[0]
		不支持python负下标取值,要取最后一个元素: lastNum := arr1[len(arr1)-1]
		3. 基于下标修改某个值 arr1[1] = 10
	*/

	var arr1 [3]int8

	// 数组赋值
	arr1 = [3]int8{0, 1}
	fmt.Println("数组是", arr1)

	// 数组修改元素
	var newNum int64 = 10000
	// int64 转 int8 时，只保留最低的 8 位二进制数据
	arr1[2] = int8(newNum)
	fmt.Println("新数组:", arr1)
}

// 在原始切片中修改切片中所有满足条件的值
func UpdateAll[T comparable](slice1 []T, oldValu, newValue T) (err error, result []T) {
	if len(slice1) == 0 {
		return fmt.Errorf("slice1不能为空"), result
	}
	// 异常恢复使用defer + recover(), defer没有返回值
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Sprintf("upodateAll fail, fail is %s", err)
		}
	}()

	for index, value := range slice1 {
		if value == oldValu {
			slice1[index] = newValue
		}
	}
	return nil, slice1

}

// 在原始切片中修改切片中满足条件的第一个值
func UpdateFirst[T comparable](slice1 []T, oldValue, newValue T) (err error, result []T) {
	if len(slice1) == 0 {
		return fmt.Errorf("slice1不能为空"), result
	}

	defer func() {
		if err := recover(); err != nil {
			err = fmt.Sprintf("修改某一个错误，错误是%s\n", err)
		}
	}()

	for index, value := range slice1 {
		if value == oldValue {
			slice1[index] = newValue
			break
		}
	}
	return nil, result
}

// 在原始切片中修改切片中满足条件的最后一个值
func UpdateLast[T comparable](slices1 []T, oldValue, newValue T) (err error, result []T) {
	sliceLen := len(slices1)

	if sliceLen == 0 {
		return fmt.Errorf("slice1不能为0"), result
	}

	defer func() {
		if err := recover(); err != nil {
			err = fmt.Sprintf("修改最后一个参数错误,错误是%s\n", err)
		}
	}()

	for i := sliceLen - 1; i >= 0; i-- {
		if slices1[i] == oldValue {
			slices1[i] = newValue
			break
		}
	}

	return nil, slices1

}

// 在原始切片中删除切片中的所有满足条件的元素
func DeleteFun[T comparable](slices1 []T, oldValues ...T) (err error, result []T) {

	sliceLen := len(slices1)

	if sliceLen == 0 {
		return fmt.Errorf("slice1不能为0"), result
	}

	defer func() {
		if err := recover(); err != nil {
			err = fmt.Sprintf("修改最后一个参数错误,错误是%s\n", err)
		}
	}()

	for _, oldValue := range oldValues {
		i := 0
		for _, value := range slices1 {
			if value != oldValue {
				slices1[i] = value
				i++
			}
		}
	}

	return nil, slices1

}

// slice切片
func SliceFun() {
	/*
		slice是一个不定长的数组，底层是一个包含指向底层数组的指针,长度和容量的结构体
		slice类型表示 []T
		slice声明如果需要使用，需要提前make([]T, len， cap)分配内存空间
		slice赋值、取值、修改值和数组一致
		slice扩容规则:
			1. 容量 < 1024：每次翻倍
			2. 容量 ≥ 1024：每次增加 25%
		slice删除元素需要for遍历原slice中的元素，匹配要删除的值，基于值对应的下标在做切片，
		如元素在下标为2的位置，则slice1[:2], 最后使用slice1 = append(slice1[:2], slice1[3:])得到删除后slice
	*/

	// 1. 声明一个切片
	var slice1 []int8 = make([]int8, 10)
	// 2. 切片赋值
	slice1 = []int8{1, 3, 4, 2, 3, 4, 5, 2, 1, 9, 10}
	fmt.Println("slice1:", slice1)
	fmt.Printf("slice1地址是%p\n:", &slice1)

	// 冒等式声明并赋值一个切片
	slice2 := []string{"hello", "hello", "heel0", "hell0"}
	fmt.Println("slice2:", slice2)

	// 基于下标切片修改值
	slice1[0] = 10
	fmt.Println("修改第一个元素值后slice1:", slice1)

	// 切片中追加元素
	slice1 = append(slice1, 110)
	fmt.Println("追加一个元素值后slice1:", slice1)
	fmt.Printf("追加一个元素值后slice1地址是%p\n:", &slice1)

	// 追加其他切片, 容量<1024，每次扩容追加1024，容量>1024，每次追加25%
	/// append(slice1, []int8{1, 2, 3}...)中 ... 是切片展开，等于逐个添加
	slice1 = append(slice1, []int8{1, 2, 3}...)
	fmt.Println("追加多个值后slice1:", slice1)
	fmt.Printf("追加多个元素值后slice1地址是%p\n:", &slice1)

}

// map映射
func MapFun() {
	/*
		map是一个包含多个键值对的数据集合，类型表示 map[string|int]T
			key必须是唯一的，可以是数字和字符串
			value可以是任意数据类型包括函数/接口/结构体等
		map声明后需要 make() 分配空间
	*/
	// 1.1 声明并分配空间给一个map
	var map1 map[float64]int = make(map[float64]int, 3)
	// 1.2 赋值
	map1 = map[float64]int{3.14: 3, 365: 365}
	fmt.Println("map1:", map1)

	// 2. 冒等式声明并赋值一个map
	map2 := map[bool]int{true: 1, false: 0}
	fmt.Println("map2:", map2)

	// map修改value
	map1[3.14] = 5
	fmt.Println("修改后key3.14的map1:", map1)

	// map修改不存在的key的值，会做添加
	map1[12] = 12
	fmt.Println("修改key12的map1:", map1)

	// 获取不存在的值,会返回value对应类型的默认值，如value是string，默认值是"""
	fmt.Printf("获取map1中key为13的值:%#v\n", map1[13])

	// 删除map中的某个key
	delete(map2, false)

}
