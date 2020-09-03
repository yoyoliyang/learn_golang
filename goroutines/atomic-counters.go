package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		// 创建50个协程
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				//每个协程原子计数增加1,遍历1000次
				// 使用原子计数器来增加ops计数，使协程之间不会互相干扰
				atomic.AddUint64(&ops, 1)

			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
