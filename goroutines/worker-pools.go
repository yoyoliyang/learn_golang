package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "started job ", j)
		// 发送jobs数据给result作为结果
		result <- j * 2
	}
}

func main() {

	const numJobs = 5
	// 初始化channel和容量
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 5; w++ {
		// 创建5个协程
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
