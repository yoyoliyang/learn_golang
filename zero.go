package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	// var s string
	const s string = "hello"
	fmt.Printf("%v,%v,%v,%q", i, f, b, s)
	// 0,0,false,""
}
