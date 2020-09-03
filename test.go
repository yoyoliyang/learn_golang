package main

import "fmt"

type IPaddr [4]byte
type Text []string

func main() {
	s0 := Text{"hello,world", "oceanlee"}
	s := IPaddr{127, 0, 0, 1}
	fmt.Println(s, s0)
}
