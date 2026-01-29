package fileoperate

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadAllFile() (err error, result string) {
	// 打开文件
	filed, err := os.Open("../main.go")
	// 判断是否有错误
	if err != nil {
		return fmt.Errorf("打开文件错误, 错误是%s", err), ""
	}

	// 一定记得关闭文件
	defer filed.Close()

	// 一次性读取文件所有信息
	bytes, err := io.ReadAll(filed)
	if err != nil {
		return fmt.Errorf("读取文件错误, 错误是%s", err), ""
	}

	return nil, string(bytes)
}

func ReadFile() (err error, result string) {
	// 打开文件
	cwdPath, _ := os.Getwd()
	fmt.Println("当前目录是", cwdPath)

	filed, err := os.Open("../main.go")
	if err != nil {
		return fmt.Errorf("打开文件错误，错误是%s", err), ""
	}

	// 每次读取固定大小
	readBytes := make([]byte, 100)
	buffers := bufio.NewReader(filed)
	for {
		n, err := buffers.Read(readBytes)
		if err != nil && err != io.EOF {
			return fmt.Errorf("读取错误, 错误是%s", err), ""
		}

		if err == io.EOF || n == 0 {
			// 读取完毕
			break
		}

	}

	return nil, string(readBytes)
}
