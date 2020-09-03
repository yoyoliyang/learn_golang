package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		// select channel选择器
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	c3 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c3 <- "result c3"
	}()

	select {
	case msg := <-c3:
		fmt.Println(msg)
		break
	case <-time.After(1 * time.Second):
		// time.After返回一个通道,在该时间前c3未完成，所以选择该case
		fmt.Println("c3 was late")
		break
	}

	c4 := make(chan string, 1)
	c4 <- "hello"
	select {
	case msg := <-c4:
		fmt.Println("c4", msg)
	default:
		// c4没数据的情况下，返回该default
		fmt.Println("c4没数据")
	}
}
