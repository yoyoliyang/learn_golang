package main

import (
	"fmt"
	"time"
)

// chan 读写通道
// chan<- 只写通道
// <-chan只读通道
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")

	done <- true
}

// 只写通道参数ping
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 只读通道参数ping
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	done := make(chan bool, 1)
	go worker(done)

	// 等待通道结束
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
