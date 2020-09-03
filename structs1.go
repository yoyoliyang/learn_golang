package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

// 指针做欸函数方法的调用对象(高效,因为直接使用会复制该对象值）
func (p *Person) talk() string {
	p.name = "li"
	return fmt.Sprintf("My name is %v, I'm %v years old", p.name, p.age)
}

func main() {

	tom := Person{
		"tom",
		6,
	}

	fmt.Println(tom.talk())

}
