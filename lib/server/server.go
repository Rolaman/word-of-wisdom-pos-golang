package server

import (
	"log"
	"net"
	"time"
)

type Server struct {
	address string
	service *BookService
}

func NewServer(address string, service *BookService) *Server {
	return &Server{
		address: address,
		service: service,
	}
}

func (s *Server) Start() error {
	socket, err := net.Listen("tcp", s.address)
	defer socket.Close()
	if err != nil {
		return err
	}
	for {
		con, err := socket.Accept()
		if err != nil {
			log.Printf("can't set connection: %v", con)
			continue
		}
		go s.process(con)
	}
}

func (s *Server) process(con net.Conn) {
	defer con.Close()
	if err := con.SetDeadline(time.Now().Add(60 * time.Second)); err != nil {
		log.Printf("can't set deadline: %v", err)
		return
	}
	if err := s.service.HandleRequest(con); err != nil {
		log.Printf("error while processing request: %v", err)
	}
}
