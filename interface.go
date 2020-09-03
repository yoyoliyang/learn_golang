package main

import "fmt"

type Name interface {
	talk() string
}

type Person struct {
	name string
}

func (p *Person) talk() string {
	return p.name
}

func main() {
	n := &Person{"liyang"}
	fmt.Println(n.talk())

}
