package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {

	defer conn.Close()

	fmt.Println("client connected:", conn.RemoteAddr())

	// 接受数据的缓冲区
	buf := make([]byte, 1024)

	for {
		// n 表示实际接受到的字节数
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println("read error:", err)
			}
			return
		}
		fmt.Println("received:", string(buf[:n]))

		// 回显数据
		if _, err := conn.Write(buf[:n]); err != nil {
			log.Println("write error:", err)
			return
		}
	}
}

func main() {

	// 使用 tcp 协议，监听本机的 8080 端口
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("tcp server listening on :8080")

	for {
		// 接受客户段的连接，Accept() 是阻塞方法，没有新的连接时一直等待
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		// 开启一个协程来处理一个客户端的连接，服务端可以同时处理多个连接
		// 不使用协程的话，那么只能一个一个处理客户段的连接
		go handleConn(conn)
	}
}
