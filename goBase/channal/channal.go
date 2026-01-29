package channal

import (
	"fmt"
	"sync"
)

func ReadGoroutine(waitGroup *sync.WaitGroup, ch <-chan int) {
	// 协程退出时waitGroup减1
	defer waitGroup.Done()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("chan读取异常, 异常是", err)
		}
	}()

	for {
		select {
		case getChan := <-ch:
			fmt.Println("从其他协程中获取的信息是", getChan)
		default:
			fmt.Println("读取完毕")
			break
		}
	}
}

func WriteGoroutine(waitGroup *sync.WaitGroup, ch chan<- int) {
	defer waitGroup.Done()

	// 捕获错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("chan写入失败, 错误是", err)
		}
	}()

	for i := 0; i < 30; i++ {
		ch <- i
	}
}
