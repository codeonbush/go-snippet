/**
  tcp粘包问题验证
*/
package main

import (
	"fmt"
	"net"
)

import (
	"strconv"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	request := make([]byte, 128)
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		if read_len == 0 {
			break
		} else {
			fmt.Println("Received data: " + string(request[:read_len]))
		}
	}
}

func startClient() {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()

	data := "Hello World!"
	data += strings.Repeat("A", 128-len(data))

	for i := 0; i < 10000; i++ {
		message := data + strconv.Itoa(i) + "\n"
		conn.Write([]byte(message))
		fmt.Println("Sent data: " + message)
		//客户端发送数据速度较快时发生粘包现象
		//time.Sleep(10 * time.Millisecond)
	}
}

func init() {
	go startClient()
}
