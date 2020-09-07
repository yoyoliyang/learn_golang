package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "hello 世界 !"
	// 将字符串转化成数组(根据空格)
	fmt.Println(strings.Fields(str))
	s := wordCount(str)
	for str, item := range s {
		fmt.Println(str, item)
	}
	fmt.Println("-------------")

	for _, item := range str {
		// 字符串遍历方法%c
		fmt.Printf("%c\n", item)
	}

	// str1 := []byte("hello,世界")
	str1 := "hello,世界"

	// fmt.Printf("%v\n", str1)
	fmt.Printf("%+q\n", str1) // 中文utf8解码,识别中文切片长度
	fmt.Printf("%U\n", str1)  // 中文utf8解码,识别中文切片长度

	// r, size := utf8.DecodeLastRune(str1)
	// 使用utf8函数来对字符串进行切片 具体看https://blog.golang.org/strings
	fmt.Println(len(str1))
	str1 = "世界"
	r, size := utf8.DecodeRuneInString(str1[3:])
	fmt.Printf("%c %v\n", r, size)

	fmt.Println("-------------")
}

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, item := range strings.Fields(s) {
		m[item] = len(item)
	}
	return m
}
