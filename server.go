package httpserver

import (
	"fmt"
	"net"
)

type HTTPServer interface {
	net.Listener
}

type HTTPServerListenError error

func NewServer() (HTTPServer, error) {
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return nil, fmt.Errorf("http server listen error, %v", err)
	}
	return server, nil
}
