package httpserver

import (
	"fmt"
	"io"
	"net"
	"os"
)

type HTTPServer interface {
	net.Listener
}

type HTTPServerListenError error

func handleConnection(conn net.Conn) {
	io.Copy(os.Stdout, conn)
}

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
