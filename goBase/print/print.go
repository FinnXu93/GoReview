// 声明包，同一文件夹下的所有go文件都应该用同一个包名，包名最好和文件夹名一致
package print

// 导入需要的包，多个包放在一个 () 中
import (
	"fmt"
	"strings"
)

// go变量/函数名首字母区分大小写，小写的只能在本界面使用，大写才能被其他包调用
func PrintFun() {
	// Println 字符串间自动添加空格,会自动换行
	var s1 string = "Print相关方法"
	fmt.Println("println输出的是", s1)

	// Printf 用占位符代表变量，格式化输出，不会自动换行
	fmt.Printf("Printf输出的是%s\n", s1)

	// Print 需要指定输出端, 不会自动换行
	fmt.Print("Print输出的是", s1)
	fmt.Println("println输出的是", s1)

	fmt.Println("---------------------------")

	/*
		Sprintf()拼接后的结果是 Sprintf输出的是Print相关方法
		最后一个字符是true

		Sprintln拼接后的结果是 Sprintln输出的是 Print相关方法
		最后一个字符是false

		Sprint拼接后的结果是 Sprint输出的是Print相关方法
		最后一个字符是true
	*/

	// Sprintf(): 用占位符拼接字符串，返回一个字符窜,也可以用于其他类型转string
	var s2 string = fmt.Sprintf("Sprintf输出的是%s", s1)
	fmt.Println("Sprintf()拼接后的结果是", s2)
	fmt.Printf("最后一个字符是%t\n", strings.HasSuffix(s2, "法"))

	var s3 string = fmt.Sprintln("Sprintln输出的是", s1)
	fmt.Println("Sprintln拼接后的结果是", s3)
	fmt.Printf("最后一个字符是%t\n", strings.HasSuffix(s3, "法"))

	var s4 string = fmt.Sprint("Sprint输出的是", s1)
	fmt.Println("Sprint拼接后的结果是", s4)
	fmt.Printf("最后一个字符是%t\n", strings.HasSuffix(s4, "法"))

	fmt.Println("---------------------------")

}
