package main

import "fmt"

func main() {
	fb(10)
}

func fb(n int) {
	pre := 0
	next := 1
	var result int

	for i := 0; i < n; i++ {
		if n == 1 {
			fmt.Println(n)
		} else {
			result = pre + next
			pre = next
			next = result
			//Printf函数 注意类型
			fmt.Printf("%T\t%v\n", result, result)
		}
	}

}
