package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/wookesh/bazel-demo/echo/proto"
	"github.com/wookesh/bazel-demo/echo/service/server"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 7000, "Server port")
)

func main() {
	s := grpc.NewServer()
	es := server.NewEchoServer()

	proto.RegisterEchoServer(s, es)

	addr := fmt.Sprintf("0.0.0.0:%v", *port)

	l, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalf("net.Listen() failed: %v", err)
	}
	log.Printf("Starting server at %v", addr)
	if err := s.Serve(l); err != nil {
		log.Fatalf("s.Serve() failed: %v", err)
	}
}
