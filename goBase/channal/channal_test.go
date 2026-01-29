package channal

import (
	"sync"
	"testing"
)

func TestChannal(t *testing.T) {
	var waitGroup sync.WaitGroup

	// 有缓存的chan
	var ch chan int = make(chan int, 10)

	defer close(ch)

	waitGroup.Add(1)
	go WriteGoroutine(&waitGroup, ch)

	waitGroup.Add(1)
	go ReadGoroutine(&waitGroup, ch)

	waitGroup.Wait()
}
