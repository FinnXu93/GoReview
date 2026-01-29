package server

import (
	"fmt"
	"net"
	"time"
)

func oneConn(conn net.Conn) {
	fmt.Printf("客户端%s已上线\n", conn.LocalAddr().Network())
	fmt.Println("欢迎光临！")
	defer conn.Close()

	var recvInfo []byte
	for {
		n, err := conn.Read(recvInfo)
		if err != nil {
			fmt.Println("服务端接受失败")
			break
		}
		fmt.Println("服务端接受到的客户端ip: ", conn.RemoteAddr(),
			"的信息: ", string(recvInfo[:n]))

		time.Sleep(5 * time.Millisecond)

		n, wErr := conn.Write(recvInfo)
		if wErr != nil {
			fmt.Println("服务端发送错误, 错误是", wErr)
			break
		}
	}

}

func CreateServer() {
	// 创建服务器连接
	listens, err := net.ListenTCP("tcp",
		&net.TCPAddr{Port: 8089})
	if err != nil {
		fmt.Println("服务器启动失败, 错误是", err)
		return
	}

	defer listens.Close()

	// 等待接受客户端连接
	for {
		conn, err := listens.Accept()
		if err != nil {
			fmt.Println("客户端连接失败, 客户端ip是", conn.RemoteAddr(), "错误是", err)
			continue
		}

		fmt.Println("客户端 ", conn.RemoteAddr(), "连接成功!")

		go oneConn(conn)
	}

}
