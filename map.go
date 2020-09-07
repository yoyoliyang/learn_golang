package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

var m map[string]Person

func main() {

	// 初始化map, 结构体map
	m = make(map[string]Person)

	m["ocean"] = Person{
		"oceanlee",
		22,
	}
	fmt.Println(m["ocean"].name)

	// 结构体map赋值方式
	n := map[string]Person{
		"liyang": Person{
			"liyang",
			22,
		},
		"oceanlee": Person{
			"oceanlee",
			22,
		},
	}
	fmt.Println(n["oceanlee"].name)

	// 整型map
	s := map[string]int{
		"number_one": 10,
		"number_two": 20,
	}

	fmt.Println(s["number_one"])
	s["number_three"] = 30
	fmt.Println(s["number_three"])
	delete(s, "number_three")
	_, ok := s["number_three"]
	fmt.Println(s["number_three"], ok)

	x := make(map[int]string)
	var l [10]int
	for i := range l {
		fmt.Println(i)
		x[i] = fmt.Sprintf("number %d", i)
	}
	fmt.Println(x)
	z := &x // 通过引用的方式获取x数据（防止复制产生的性能损失）

	// 指针迭代方式和指针索引映射数据
	for index, item := range *z {
		(*z)[index] = fmt.Sprintln(index + 1)
		fmt.Println((*z)[index], item)
	}

	map1 := make(map[int]string)
	map1[0] = "青花瓷"
	// 判断是否存在映射
	el, ok := map1[1]
	if ok {
		fmt.Println(el, ok)
	} else {
		fmt.Println(ok)
	}

}
