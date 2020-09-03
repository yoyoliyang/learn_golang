package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)
	// 初始化一个切片
	var list1 = make([]int, 10)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOPs uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			total1 := 0
			for {
				key := rand.Intn(5)
				// 锁定state和list1
				mutex.Lock()
				// 读取state(使用)
				total += state[key]
				total1 += list1[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				key_list1 := rand.Intn(10)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				list1[key_list1] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOPs, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOPs)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state", state)
	fmt.Println("list1", list1)
	mutex.Unlock()

}
