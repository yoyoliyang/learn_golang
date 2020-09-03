package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 利用协程批量制作随机整数
	result := make(chan int, 10)
	for i := 0; i < 10; i++ {

		go func() {
			key := rand.Intn(50)
			fmt.Println(key)
			result <- key
		}()
	}

	for i := 0; i < 10; i++ {
		<-result
	}
}
