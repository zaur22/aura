package rpc

import (
	"context"
	"fmt"
	"github.com/oklog/run"
	"github.com/zaur22/aura/pkg/rpc/api"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Run(ctx context.Context, addr string) error
}

func NewServer(dto NewServerDTO) Server {
	s := grpc.NewServer()
	incrServ := newIncrementerServer(dto.Incrementer)
	api.RegisterIncrementerServer(s, incrServ)
	return &server{
		server: s,
	}
}

type server struct {
	server *grpc.Server
}

func (s *server) Run(ctx context.Context, addr string) error {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Did not connect: %v", err)
	}
	defer conn.Close()

	var g run.Group
	{
		g.Add(func() error {
			return s.server.Serve(conn)
		},
			func(err error) {
				s.server.GracefulStop()
			})
	}
	{
		var stop = make(chan struct{})
		g.Add(func() error {
			select {
			case <-ctx.Done():
			case <-stop:
			}
			return nil
		},
			func(err error) {
				close(stop)
			})
	}

	return g.Run()
}
