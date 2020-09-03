package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// i := 0
	// for i < 10 {
	// fmt.Println(i)
	// i++
	s1 := []int{11, 2, 3, 4, 5}
	for index, item := range s1 {
		fmt.Printf("%v\t%v\n", index, item)
	}
	for _, item := range s1 {
		fmt.Println(item)
	}

}
