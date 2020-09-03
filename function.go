package main

import "fmt"

//函数作为参数
func func1(fn func(int, int) int) {
	fmt.Println(fn(10, 10))
}

// 闭包函数
func func3() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println("i =", i)
		return i
	}

}

func func2(fn func(s string) string) {
	fmt.Println(fn("world!"))
}

// 变参函数（任意数目的参数）
func sum(nums ...int) int {
	s := 0
	for i := range nums {
		fmt.Println(i)
		s += i
	}
	return s

}

func main() {

	//  函数传递给变量
	test := func(x int, y int) int {
		return x + y
	}

	test1 := func(x int, y int) int {
		fmt.Printf("hello, x + y = %d\n", x+y)
		return x + y
	}

	test2 := func(s string) string {
		return fmt.Sprintf("hello %q", s)
	}

	func1(test)
	func1(test1)
	func2(test2)

	// 直接运行匿名函数
	func() string {
		fmt.Println("world")
		return "world"
	}() // 注意此处括号

	// 闭包函数
	test3 := func3()          // func3()返回的是一个函数,test3，执行该函数需要再次()
	fmt.Println("1", test3()) // test3 = 1
	fmt.Println("2", test3())
	fmt.Println("3", test3())

	l1 := []int{1, 2, 3, 4, 5}
	// 切片作为变参函数的调用方法
	fmt.Println(sum(l1...))
	fmt.Println(sum(1, 2, 3, 4, 5))

}
