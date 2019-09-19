package main

import (
	"fmt"
	"net"
	"strings"
	"bytes"
)

func connHandler(c net.Conn) {
	if c == nil {
		return
	}

	buf := make([]byte, 4096)
	var buf_msg bytes.Buffer

	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}

		input_string := strings.TrimSpace(string(buf[:cnt]))
		inputs := strings.Split(input_string, " ")

		switch inputs[0] {
		// sys cmd
		case "quit":
			buf_msg.Reset()
			buf_msg.WriteString("Bye Bye...\n")
			c.Write([]byte(buf_msg.String()))
			c.Close()
			break
		default:
			buf_msg.Reset()
			buf_msg.WriteString(">> ")
			buf_msg.WriteString(input_string)
			buf_msg.WriteString("\n")
			c.Write([]byte(buf_msg.String()))
		}
	}

	fmt.Println("Connection from", c.RemoteAddr(), "closed")
}

func main() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
		return
	}

	fmt.Println("Server Started ...")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		
		go connHandler(conn)
	}
}
