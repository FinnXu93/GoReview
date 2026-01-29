package client

import (
	"fmt"
	"net"
	"sync"
)

func clientConn(conn net.Conn, waitGroup *sync.WaitGroup, data string) {
	defer conn.Close()
	defer waitGroup.Done()

	for {
		_, wErr := conn.Write([]byte(data))
		if wErr != nil {
			fmt.Println("客户端发送失败, 错误是", wErr)
			break
		}
	}
}

func CreateClient(addr string, data string) {
	conn, err := net.DialTCP("tcp", &net.TCPAddr{Port: 8085}, &net.TCPAddr{Port: 8085})
	if err != nil {
		fmt.Println("客户端连接错误, 错误是", err)
		defer conn.Close()
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go clientConn(conn, &waitGroup, data)
	waitGroup.Wait()
}
