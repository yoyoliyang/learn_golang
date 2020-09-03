package main

import "fmt"

const (
	BIG = 1 << 10
	// 二进制表示方式 1后面10个0
	SMALL = BIG >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(BIG)
	// fmt.Println(needFloat(SMALL))
	// fmt.Println(needFloat(BIG))
}
