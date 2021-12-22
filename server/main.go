package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zhou-en/grpc_tutorials/proto"
)

type server struct {
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Add(context.Context, *Request) (*Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(context.Context, *Request) (*Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
