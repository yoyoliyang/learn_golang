package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	// done用来通知main,提示任务关闭
	done := make(chan bool, 1)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j < 5; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	// 关闭通道
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	j1 := make(chan string, 2)
	j1 <- "hello"
	j1 <- "world"
	close(j1)
	for i := range j1 {
		fmt.Println(i)
	}

}
