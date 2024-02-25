package option

import "fmt"

type Option func(*Server) error

type Server struct {
	addr    string
	port    int
	maxConn int
}

func NewServer(addr string, opts ...Option) (*Server, error) {
	svc := &Server{addr: addr}
	fmt.Printf("len(opts): %v\n", len(opts))
	for _, o := range opts {
		if err := o(svc); err != nil {
			return nil, err
		}
	}
	return svc, nil
}

func WithPort(port int) Option {
	return func(s *Server) error {
		if port < 0 || port > 65535 {
			return fmt.Errorf("port must be between 0 and 65535")
		}
		s.port = port
		return nil
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) error {
		if maxConn < 0 || maxConn > 10000 {
			return fmt.Errorf("maxConn must be between 0 and 10000")
		}
		s.maxConn = maxConn
		return nil
	}
}
