package server

import (
	"context"

	"echo/proto"
)

type echoServer struct{}

func NewEchoServer() *echoServer {
	return &echoServer{}
}

func (e *echoServer) Ping(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	return nil, nil
}
