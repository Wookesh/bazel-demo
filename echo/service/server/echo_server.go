package server

import (
	"context"

	"echo/proto"
)

type echoServer struct {
	proto.UnimplementedEchoServer
}

func NewEchoServer() *echoServer {
	return &echoServer{}
}

func (e *echoServer) Ping(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	return &proto.EchoResponse{Message: req.GetMessage()}, nil
}
