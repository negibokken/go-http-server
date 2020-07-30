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

type Tokenizer struct {
	scanner *bufio.Scanner
}

// Response = Status-Line      ; Section 6.1
// 					*(( general-header ; Section 4.5
// 					| response-header  ; Section 6.2
// 					| entity-header ) CRLF) ; Section 7.1
// 					CRLF
// 					[ message-body ] ; Section 7.2
// Status-Line
// Status-Line = HTTP-Version SP Status-Code SP Reason-Phrase CRLF

func (t Tokenizer) nextChar() string {
	t.scanner.Scan()
	return t.scanner.Text()
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanRunes)
	tokenizer := &Tokenizer{scanner}
	for {
		c := tokenizer.nextChar()
		fmt.Printf("%s\n", c)
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
