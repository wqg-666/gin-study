package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	message := "Hello wqg"

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server response:", string(buf[:n]))
	//server response: Hello TCP Server
}
