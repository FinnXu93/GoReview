package pathoperate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
os.Stat(path): 可以获取大小，是否是目录，名称等
path.filepath包:
filepath.base(): .
filepath.Abs(): /Users/finnxu/Projects/GoProjects/GoReview/goBase/pathOperate
filepath.Rel(): ..
filepath.Dir(): .
filepath.SpliteList(): [./]
filepath.Join(): main.go
filepath.Glob(): [pathOperate.go pathOperate_test.go]
*/

func FilePathOperate(p string) {
	// 路径名和路径操作
	// 获取路径信息中最后一段信息
	basePash := filepath.Base(p)
	fmt.Println("filepath.base():", basePash)

	// 获取绝对路径可能会失败
	adsPath, err := filepath.Abs(p)
	if err != nil {
		fmt.Println("filepath.Abs()错误, 错误是", err)
		return
	}
	fmt.Println("filepath.Abs():", adsPath)

	// 获取和目标路径的相对路径
	relPath, err := filepath.Rel(p, "../")
	if err != nil {
		fmt.Println("filepath.Rel()错误, 错误是", err)
		return
	}
	fmt.Println("filepath.Rel():", relPath)

	// 获取目录
	dir := filepath.Dir(p)
	fmt.Println("filepath.Dir():", dir)

	// 路径分割
	paths := filepath.SplitList(p)
	fmt.Println("filepath.SpliteList():", paths)

	// 路径拼接
	newPath := filepath.Join(p, "main.go")
	fmt.Println("filepath.Join():", newPath)

	// 文件匹配
	matchs, err := filepath.Glob("./*.go")
	if err != nil {
		fmt.Println("filepath.Glob()错误, 错误是", err)
		return
	}
	fmt.Println("filepath.Glob():", matchs)

	// 获取文件扩展名
	ext := filepath.Ext(p)
	fmt.Println("filepath.Ext():", ext)

}

func DirWalk(p string) {
	// 目录遍历
	err := filepath.WalkDir(p, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("WalkDir内部错误, 错误是", err)
			return err
		}
		dirInfo, err := info.Info()
		if err != nil {
			fmt.Println("info.Info()错误, 错误是", err)
			return err
		}
		name := dirInfo.Name()
		fmt.Println("每个文件或目录的名称:", name)
		isDir := dirInfo.IsDir()
		fmt.Println("判断每个文件是否是目录:", isDir)
		size := dirInfo.Size()
		fmt.Println("判断每个文件或目录的大小:", size)

		return nil
	})

	if err != nil {
		fmt.Println("filepath.WalkDir错误, 错误是", err)
	}
}

func StatOperate(path string) {
	// os.Stat()操作
	// 默认是linux系统判断路径是否合法
	trimSpacePath := strings.TrimSpace(path)
	// 字符串切割
	stringSlice := strings.Split(trimSpacePath, "/")
	fmt.Println("path以/切割后的结果是", stringSlice)

	switch {
	case !strings.Contains(trimSpacePath, "/"):
		fmt.Println("路径中不包括 / ")
		return
	case len(stringSlice) < 2:
		fmt.Println("路径不符合要求")
		return
	}

	// 判断目录是否存在
	info, err := os.Stat(trimSpacePath)
	if err != nil {
		fmt.Println("获取目录信息异常, 异常是", err)
		return
	}

	name := info.Name()
	fmt.Println("文件或目录名是", name)

	size := info.Size()
	fmt.Println("文件或目录的大小是", size)

	isDir := info.IsDir()
	fmt.Println("path对应的是否目录:", isDir)

	if isDir {
		files, err := os.ReadDir(trimSpacePath)
		if err != nil {
			fmt.Println("读取目录错误, 错误是", err)
			return
		}
		fmt.Println("读取目录信息是", files)
	}

}

func DirOperate(DirPath string) {
	// 目录操作
	// 判断目录是否存在
	fmt.Println("os.MkdirAll()当目录结构不存在时创建目录结构.")
	err := os.MkdirAll(DirPath, 655)
	fmt.Println("os.MkdirAll()错误, 错误是", err)
}
