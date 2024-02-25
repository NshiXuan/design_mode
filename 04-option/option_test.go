package option

import (
	"fmt"
	"testing"
)

func TestOption(t *testing.T) {
	server, err := NewServer("127.0.0.1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("server: %v\n", server)
	server, err = NewServer("127.0.0.1", WithPort(8080))
	if err != nil {
		panic(err)
	}
	fmt.Printf("server: %v\n", server)
	server, err = NewServer("127.0.0.1", WithPort(8080), WithMaxConn(100))
	if err != nil {
		panic(err)
	}
	fmt.Printf("server: %v\n", server)

	server, err = NewServer("127.0.0.1", WithPort(80808))
	if err != nil {
		panic(err)
	}
	fmt.Printf("server: %v\n", server)
}
