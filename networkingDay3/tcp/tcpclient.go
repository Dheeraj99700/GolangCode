//client socket

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, server!I am Client 1"))
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received message: ", string(buf[:n]))
}
