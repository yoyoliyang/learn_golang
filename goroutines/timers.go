package main

import (
	"fmt"
	"time"
)

func main() {
	// 一个定时器
	timer1 := time.NewTimer(2 * time.Second)

	// 阻塞直到打点器结束
	<-timer1.C
	fmt.Println("end timer1")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 end")
	}()

	// 等待2s，使timer2完成
	// time.Sleep(2 * time.Second)

	stop2 := timer2.Stop() // 触发前直接取消定时器timer2
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
