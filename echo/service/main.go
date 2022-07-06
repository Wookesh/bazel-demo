package main

import (
	"echo/service/server"
)

func main() {
	s := server.NewEchoServer()
	_ = s
}
