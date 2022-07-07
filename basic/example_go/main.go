package main

import (
	"flag"

	"github_com_wookesh_bazel_demo/basic/example_go/greeter"

	"github.com/sirupsen/logrus"
)

var (
	to = flag.String("to", "world", "hello to whom?")
)

func main() {
	flag.Parse()
	logrus.Infof("hello %v", *to)
	g := greeter.NewGreeter("pepe")
	logrus.Infof(g.Greet())
}
