//client socket

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, server!Client"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
