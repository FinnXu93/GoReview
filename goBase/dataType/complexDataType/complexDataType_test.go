package complexdatatype

import (
	"fmt"
	"testing"
)

func TestArrayFun(t *testing.T) {
	t.Run("arrayFun", func(t *testing.T) {
		ArrayFun()
	})
}

func TestUpdateAll(t *testing.T) {
	t.Run("updateAll", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 10, 10, 15, 10, 16, 17}
		value := 10
		newValue := 100
		err, newSlice := UpdateAll(slice1, value, newValue)
		if err == nil {
			fmt.Println("修改后的slice1是", newSlice)
			fmt.Printf("修改后的slice1的指向数组的内存地址是%p\n", *(&newSlice))
			fmt.Println("原始slice1是", slice1)
			fmt.Printf("原始slice1的指向数组的内存地址是%p\n", *(&slice1))

		} else {
			fmt.Println("修改slice1失败, 错误是", err)
		}
	})
}

func TestUpdateFirst(t *testing.T) {
	t.Run("updateFirst", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 10, 10, 15, 10, 16, 17}
		value := 10
		newValue := 100
		err, newSlice := UpdateFirst(slice1, value, newValue)
		if err == nil {
			fmt.Println("修改后的slice1是", newSlice)
			fmt.Println("原始slice1是", slice1)
		} else {
			fmt.Println("修改slice1失败, 错误是", err)
		}
	})
}

func TestUpdateLast(t *testing.T) {
	t.Run("updateLast", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 10, 10, 15, 10, 16, 17}
		value := 10
		newValue := 100
		err, newSlice := UpdateLast(slice1, value, newValue)
		if err == nil {
			fmt.Println("修改后的slice1是", newSlice)
			fmt.Println("原始slice1是", slice1)
		} else {
			fmt.Println("修改slice1失败, 错误是", err)
		}
	})
}
