package main

import (
	"fmt"
	"net"
	"time"
)

type Config struct {
}

type Server struct {
	timeout  time.Duration
	listener net.Listener
}

func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	srv := Server{listener: l}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func main() {
	srv, _ := NewServer("localhost")
	timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
	}

	srv2, _ := NewServer("localhost", timeout)
	fmt.Println(srv,srv2)
}
