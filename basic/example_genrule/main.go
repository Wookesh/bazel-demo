package main

import (
	"flag"
	"fmt"
)

var (
	count      = flag.Int("count", 5, "count to print")
	linePrefix = flag.String("prefix", "", "prefix for each line")
)

func main() {
	flag.Parse()
	prefix := *linePrefix
	if prefix == "" {
		prefix = "very not random line: "
	}

	for i := 0; i < *count; i++ {
		fmt.Printf("%v%v\n", prefix, i)
	}
}
