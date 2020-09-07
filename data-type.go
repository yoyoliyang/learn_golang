package main

import "fmt"

type Mouse struct {
	left     bool
	right    bool
	middle   bool
	position [2]float64
}

func main() {
	// a := "helllo,world"
	// b := 10
	// c := 1.2
	// var d float64 = 1.2
	// var e uint8 = 255
	// 数组
	var arr1 []string
	var arr2 [2]int
	// 切片
	slice3 := make([]string, 5)
	fmt.Println("-------------")
	fmt.Println(arr1, "\n", arr2, "\n", slice3[3])

	x := 10
	y := x // 复制x的值
	y = 11
	fmt.Println(x, y)

	z := &x // 指针 数据指向x的值
	*z = 20 // 解引用，x此时=20
	fmt.Println(z, x)

}
