package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type Chat struct {
	Name    string
	Message string
}

func main() {

	var username string
	fmt.Println("Enter UserName")
	fmt.Scan(&username)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {

		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		var data = Chat{username, line}
		fmt.Println(data)
		file, err := json.Marshal(data)
		_, err = conn.Write([]byte(file))
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
		fmt.Println(string(buf[:n]))

	}

}
