package main

import "fmt"

// 返回一个返回int的函数

func fibonacci() func() int {
	a :=0
	b :=1
	var c int
	func result() int {
		c = a+b
		a = b
		b = c
		return c
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

