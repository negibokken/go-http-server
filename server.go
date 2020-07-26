package httpserver

import (
	"bufio"
	"fmt"
	"net"
)

type HTTPServer interface {
	net.Listener
}

type HTTPServerListenError error

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("%v\n", scanner.Text())
		err := conn.Close()
		if err != nil {
			fmt.Println("connection close error %v", err)
		}
	}
}

// NewServer は tcp server を返す
// 終了時には Close すること
func NewServer() (HTTPServer, error) {
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return nil, fmt.Errorf("http server listen error, %v", err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			return nil, fmt.Errorf("http server accept error, %v", err)
		}
		go handleConnection(conn)
	}
	return server, nil
}
