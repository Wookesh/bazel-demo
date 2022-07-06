package main

import (
	"fmt"

	"echo/service/server"
)

func main() {
	s := server.NewEchoServer()
	_ = s
	fmt.Println("hello")
}
