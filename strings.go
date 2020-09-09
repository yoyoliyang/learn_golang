package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	s := "hello,world!"
	// 替换
	fmt.Println(strings.Replace(s, "he", "ye", 1))
	fmt.Println(s)

	// prefix 判断是否存在前缀
	fmt.Println(strings.HasPrefix("http://www.baidu.com", "http://"))
}
