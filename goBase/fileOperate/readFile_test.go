package fileoperate

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	err, s1 := ReadFile()
	if err != nil {
		fmt.Println("错误是", err)
	}
	fmt.Println("读取到的是", s1)
}
