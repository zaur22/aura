package rpc

import (
	"fmt"
	"github.com/zaur22/aura/pkg/rpc/api"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Run(addr string) error
}

func NewServer(dto NewServerDTO) Server {
	s := grpc.NewServer()
	incrServ := newIncrementerServer(dto.Incrementer)
	api.RegisterIncrementerServer(s, incrServ)
	return &server{server: s}
}

type server struct {
	server *grpc.Server
}


func (s *server) Run(addr string) (resErr error) {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Did not connect: %v", err)
	}
	defer conn.Close()

	return s.server.Serve(conn)
}