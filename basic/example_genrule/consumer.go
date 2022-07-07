package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("hello/example_genrule/list.txt")
	if err != nil {
		panic(err)
	}
	for _, l := range strings.Split(string(data), "\n") {
		fmt.Println(l)
	}
}
