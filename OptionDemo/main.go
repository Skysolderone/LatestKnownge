package main

import "fmt"

type Server struct {
	Name    string
	Age     int
	Detail  string
	Address string
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	// default
	srv := &Server{
		"test", 10, "0618", "usa",
	}
	// option
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

func WithName(name string) Option {
	return func(s *Server) {
		s.Name = name
	}
}

func main() {
	// default
	s := NewServer()
	fmt.Printf("%#v\n", s)

	// option
	s1 := NewServer(WithName("test option"))
	fmt.Printf("%#v", s1)
}
