package main

import "fmt"

func assess(b bool) {
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func main() {
	a := true // if判断只能使用布尔值
	assess(a)

}
