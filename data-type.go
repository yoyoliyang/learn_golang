package main

import "fmt"

func main() {
	// a := "helllo,world"
	// b := 10
	// c := 1.2
	// var d float64 = 1.2
	// var e uint8 = 255
	var slice1 []string
	var slice2 [2]int
	slice3 := []string{}
	fmt.Println(slice1, slice2, slice3)

	x := 10
	y := x // 复制x的值
	y = 11
	fmt.Println(x, y)

	z := &x // 指针 数据指向x的值
	*z = 20 // 解引用，x此时=20
	fmt.Println(z, x)

}
