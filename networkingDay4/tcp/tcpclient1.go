//client socket

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		time.Sleep(5 * time.Minute)
		_, err = conn.Write([]byte("Hello, server!Client"))
		if err != nil {
			fmt.Println(err)
			return
		}
		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(buffer[:n]))

	}

}
