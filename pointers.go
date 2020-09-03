package main

import "fmt"

func main() {
	var p *int
	num1 := 42
	p = &num1
	fmt.Println(p)
	fmt.Printf("%v on %v\n"+"\n", num1, p)
	fmt.Printf("%v type is %T\n", p, p)

	*p = 10
	fmt.Printf("num1 被改成了： %v\n", num1)

}
