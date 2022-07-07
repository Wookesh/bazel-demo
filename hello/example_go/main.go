package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var (
	to = flag.String("to", "world", "hello to whom?")
)

func main() {
	flag.Parse()
	logrus.Infof("hello %v", *to)
}
