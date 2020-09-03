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
func ping(ping chan<- string, msg string) {
	ping <- msg
}

// 只读通道参数ping
func pong(ping <-chan string, pongs chan<- string) {
	msg := ping
	pongs <- msg
}
func main() {
	go worker(done)

	// 等待通道结束
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(ping, pongs)
	fmt.Println(<-pongs)
}
