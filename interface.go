package main

import (
	"fmt"
	"io"
)

type Name interface {
	talk() string
	move() int
}

type Person struct {
	name string
}
type Cat struct {
	name string
}

func (p *Person) talk() string {
	return p.name
}
func (p *Person) move() int {
	return 30
}
func (c *Cat) talk() string {
	return c.name
}

func main() {
	n := &Person{"liyang"}
	c := &Cat{"tom"}
	fmt.Println(n.talk(), n.move())
	fmt.Println(c.talk())
}
