package build

import "fmt"

type Server struct {
	addr string
	port int
}

type ServerBuilder struct {
	addr string
	port int
}

func NewBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (s *ServerBuilder) SetAddr(addr string) error {
	s.addr = addr
	return nil
}

func (s *ServerBuilder) SetPort(port int) error {
	if port < 0 || port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535")
	}
	s.port = port
	return nil
}

func (s *ServerBuilder) Build() *Server {
	svc := &Server{}
	if len(s.addr) > 0 {
		svc.addr = s.addr
	} else {
		svc.addr = "0.0.0.0"
	}
	if s.port > 0 && s.port < 65535 {
		svc.port = s.port
	}
	return svc
}
