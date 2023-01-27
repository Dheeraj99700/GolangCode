package main

import (
	"fmt"
	"net"
)

/*type Chat struct {
	Name    string `json:"Name"`
	Message string `json:"Message"`
}*/

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//	fmt.Printf("Client connected: %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	//var chat []Chat

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
		result := buf[:n]
		_, err = conn.Write([]byte(result))
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
