package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	// listen on the specified port

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	// clean up the listener when done
	defer listener.Close()

	fmt.Printf("Listening on port %s...\n", listener.Addr())

	// infinite loop to accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		go handleConnection(conn)

	}

}

// handle the connection in a separate goroutine

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {

		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", bytes)

		line := fmt.Sprintf("Echo: %s", bytes)

		fmt.Printf("response: %s", line)
		_, err = conn.Write([]byte(line))

		// check for errors while writing to the connection
		if err != nil {
			fmt.Println("failed to write data, err:", err)
			return
		}
		fmt.Println("response sent successfully")

	}
}


// to write response from client side 
//  i) in windows : telnet localhost port .
//  ii) in linux : echo Hello world | nc localhost 9090