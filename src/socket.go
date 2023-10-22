package main

import (
	"errors"
	"fmt"
	"net"
)

type Socket struct {
	l net.Listener
}

func NewSocket() *Socket {
	return &Socket{
		l: nil,
	}
}

func (s *Socket) Listen(host string, port uint16) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		return err
	}

	s.l = listener

	return nil
}

func (s *Socket) Accept() (net.Conn, error) {
	if s.l == nil {
		return nil, errors.New("Accept() called while listener is nil")
	}

	return s.l.Accept()
}

func (s *Socket) Close() error {
	if s.l == nil {
		return nil
	}

	return s.l.Close()
}
