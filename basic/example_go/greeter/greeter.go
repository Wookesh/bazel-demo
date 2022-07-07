package greeter

import (
	"fmt"
)

type Greeter struct {
	who string
}

func (g *Greeter) Greet() string {
	return fmt.Sprintf("hello %v", g.who)
}
