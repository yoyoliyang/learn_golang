package main

import "fmt"

func Sqrt(x float64) float64 {
	// 牛顿法求平方根
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	Sqrt(10.0)
}
