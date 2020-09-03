package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个打点器，每隔500毫秒发送数据
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				// 阻塞500毫秒
				fmt.Println("打点", t)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("打点器结束")

	timer1 := make(chan time.Time, 3)
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// time.Tick可以无限遍历
			fmt.Println("打点可以range遍历", t)
			// 填充timer1,当填充到3个时，直到下方timer1取出完毕
			timer1 <- t
		}
	}()

	// 阻塞timer1,直到结束
	for i := 1; i <= 3; i++ {
		<-timer1
	}
}
