package main

import (
	"flag"

	"github.com/wookesh/bazel-demo/basic/example_go/greeter"

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
