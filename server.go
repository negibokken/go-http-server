package httpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type HTTPServer interface {
	net.Listener
}

type HTTPServerListenError error

func lex(scan bufio.Scanner) {
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	// scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Printf("%v\n", scanner.Bytes())
		if scanner.Err() != nil {
			log.Fatal("non-EOF error")
		}
	}
	fmt.Print("end conn")
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
