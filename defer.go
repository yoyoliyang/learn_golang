package main

import "fmt"

func main() {
	// defer后进先出
	defer fmt.Println("vim-go last") // 第一个，后出

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //最后一个，先出
	}
	fmt.Println("vim-go")
}
