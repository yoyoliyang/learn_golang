package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序,内建函数

	s1 := []int{11, 20, 3, 4, 5}

	s3 := []string{"a", "h", "g", "c", "-"}
	// 字符串是否被排序
	fmt.Println(sort.StringsAreSorted(s3))

	// 整型排序
	sort.Ints(s1)
	// 字符串排序
	sort.Strings(s3)
	fmt.Println(sort.StringsAreSorted(s3))

	fmt.Println(s1, "\n", s3)
}
