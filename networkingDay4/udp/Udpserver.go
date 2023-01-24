/*package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on a port
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Read data
		buffer := make([]byte, 4096)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the data
		go handleData(conn, addr, buffer[:n])
	}
}

func handleData(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	// Send a response
	conn.WriteToUDP([]byte("Hello, client!"), addr)

	// Print the received data
	fmt.Println(string(data))
}
*/
//sever
package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on UDP address and port
	addr, _ := net.ResolveUDPAddr("udp", ":8000")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	// Continuously read incoming data
	for {
		var buf [1024]byte
		n, _, _ := conn.ReadFromUDP(buf[:])
		fmt.Println(string(buf[:n]))
	}
}
