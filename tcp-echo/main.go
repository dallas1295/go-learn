package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// go run main.go 9090
	if len(os.Args) < 2 {
		fmt.Println("TCP running on: <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to print listener, err:", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Printf("listening on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed accept connection, err:", err)
			continue // we want an infinite loop so we don't stop the tcp connections
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn) // creates a buffered reader to have efficient smaller reads

	for {
		// Read until new line character is entered returns the data read as byte slice, it reads large chunks line by line
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			// EOF: client closes connection gracefully
			if err != io.EOF {
				// if not closed gracefully
				fmt.Println("failed to read data, err:", err)
			}
			return // end the infitinite loop if there's a problem
		}

		fmt.Printf("request: %s", bytes)

		line := fmt.Sprintf("Echo: %s", bytes)
		fmt.Printf("response: %s", line)

		connWrite, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("failed to write data, err:", err)
			return
		}
		fmt.Println(connWrite)
	}
}
