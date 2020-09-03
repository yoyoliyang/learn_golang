package main

import "fmt"

// type Name interface {
// // 接口定义了一组方法的数据签名
// talk() string
// }

type Person struct {
	name string
}

func (p *Person) talk() string {
	return p.name
}

func main() {
	n := &Person{"liyang"}
	fmt.Println(n.talk())

	var i interface{} = "hello"
	fmt.Printf("%v, %T\n", i, i)
	s, ok := i.(string)
	fmt.Println(s, ok)
	s1, ok1 := i.(bool)
	fmt.Println(s1, ok1)

	switch v := i.(type) {
	case string:
		fmt.Printf("string :%v", v)
	case bool:
		fmt.Printf("bool :%v", v)
	default:
		fmt.Printf("nothin :%v", v)
	}

}
