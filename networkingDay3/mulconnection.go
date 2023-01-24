package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	var input string
	fmt.Println("Do you want to connect to server.(Press Yes = 'Y', No = 'N')")
	fmt.Scanln(&input)
	switch input {
	case "n":
		os.Exit(0)
	case "y":
		// create listener
		ln, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println("Error creating listener:", err)
			return
		}
		defer ln.Close()

		var wg sync.WaitGroup

		// accept connections and handle them concurrently
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			wg.Add(1)
			go handleConnection(conn, &wg)
		}
		wg.Wait()
	}

}
func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		data := "Hi client"
		fmt.Println("Received:", string(buffer[:n]))

		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
