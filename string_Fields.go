package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world !"
	fmt.Println(strings.Fields(str))
	s := wordCount(str)
	for str, item := range s {
		fmt.Println(str, item)
	}
}

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, item := range strings.Fields(s) {
		m[item] = len(item)
	}
	return m
}
