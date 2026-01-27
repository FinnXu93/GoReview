package scan

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	t.Run("scanFun", func(t *testing.T) {
		fmt.Println("请输入年龄:")

		var age int8
		fmt.Scan(&age)

		fmt.Println("输入的年龄是:", age)
		// ScanFun()
	})
}
