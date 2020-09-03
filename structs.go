package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

type Rectangle struct {
	width  int
	heigth int
}

func (p Person) Speak() string {
	return fmt.Sprintf("hello, my name is %v", p.name)
}

// 上同
func Speak1(p Person, msg string) string {
	return fmt.Sprintf("%v, my name is %v", msg, p.name)
}

func main() {
	person := Person{name: "liyang", age: 33}
	p := &person
	fmt.Println(person.name, person.age)

	person.name = "oceanlee"
	fmt.Println(person.name, person.age)
	fmt.Println((*p).name, person.age) // 通过指针访问结构体数据
	fmt.Println(p.name, person.age)    // 上同，隐式引用方法

	a := Rectangle{20, 30}
	x := Rectangle{width: 10, heigth: 20} // 字段都赋值
	y := Rectangle{width: 10}             // heigth此时隐式赋值为0
	z := Rectangle{}                      // w := 0,h:=0
	p1 := &Rectangle{10, 20}
	fmt.Println(a, x, y, z, p1.width)

	fmt.Println(person.Speak())
	fmt.Println(Speak1(person, "hello"))

}
