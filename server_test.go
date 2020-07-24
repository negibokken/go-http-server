package httpserver

import "testing"

func TestNewServer(t *testing.T) {
	server, err := NewServer()
	defer server.Close()
	if err != nil {
		t.Fatal("No server here")
	}
}
