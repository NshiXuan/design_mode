package build

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	b := NewBuilder()
	if err := b.SetAddr("127.0.0.1"); err != nil {
		panic(err)
	}
	if err := b.SetPort(8080); err != nil {
		panic(err)
	}
	server := b.Build()
	fmt.Printf("server: %v\n", server)
}
