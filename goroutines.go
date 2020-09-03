package main

import (
	"fmt"
	"time"
)

func f1(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, i)
	}
}

func main() {

	// 初始化一个通道数据,类型为string
	channel_msg := make(chan string)
	f1("world")
	// go协程并发函数
	go func() {
		fmt.Println("<<<")
		channel_msg <- "协程1传递过来"
	}()
	go func() {
		fmt.Println(">>>")
	}()
	go f1("hello")
	// 通道数据传递出来
	msg := <-channel_msg
	fmt.Println(msg)
	time.Sleep(time.Second)
	fmt.Println("done")

	// 初始化一个2个值缓冲的通道
	msg1 := make(chan string, 2)
	msg1 <- "hello"
	msg1 <- "world!"
	fmt.Println(<-msg1)
	fmt.Println(<-msg1)
}
