package main

import (
	"fmt"
	"time"
)

func main() {
	// requesets := make(chan int, 5)
	// // 填充通道
	// for i := 1; i <= 5; i++ {
	// requesets <- i
	// }
	// close(requesets)

	// limiter := time.Tick(200 * time.Millisecond)

	// for req := range requesets {
	// // 遍历每个requests并定时打点
	// <-limiter
	// fmt.Println("request", req, time.Now())
	// }

	///////////////////////////////////////////////////////////////////
	// 爆发事件
	// 初始化bursty结构体通道,缓冲数量3
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		// 填充bursty通道
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
