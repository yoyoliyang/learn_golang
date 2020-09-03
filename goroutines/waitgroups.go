package main

import (
	"fmt"
	"sync"
	"time"
)

// wg需要通过指针传递,每个协程都会运行该函数
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}
func main() {
	// wg初始化，用于等待该函数启动的所有协程
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 阻塞，知道所有协程完毕
	wg.Wait()
}
