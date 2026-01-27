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

// 修改切片中所有满足条件的值
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

// 修改切片中满足条件的第一个值
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

// 修改切片中满足条件的最后一个值
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

// 删除切片中的满足条件的元素
func DeleteFun[T comparable](slices1 []T, oldValue T) (err error, result []T) {

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

	// 冒等式声明并赋值一个切片
	slice2 := []string{"hello", "hello", "heel0", "hell0"}
	fmt.Println("slice2:", slice2)

	// 基于下标切片修改值
	slice1[0] = 10
	fmt.Println("修改第一个元素值后slice1:", slice1)
}
