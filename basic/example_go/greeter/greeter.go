package greeter

import (
	"fmt"
)

type Greeter struct {
	who string
}

func NewGreeter(who string) *Greeter {
	return &Greeter{
		who: who,
	}
}

func (g *Greeter) Greet() string {
	who := g.who
	if g.who == "" {
		who = "world"
	}
	return fmt.Sprintf("hello %v", who)
}
