package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Data{"Oceanlee", "你好世界", 1294706395881547000}
	b, err := json.Marshal(m)
	// b为字节
	fmt.Println(string(b), err)

	// 序列化
	sl1 := []string{"hello", "world"}
	mp1 := make(map[string]string)
	mp1 = map[string]string{
		"hello": "world!",
		"world": "hello!",
	}
	boolB, _ := json.Marshal(sl1)
	mpB, _ := json.Marshal(mp1)
	fmt.Println(string(boolB), err)
	fmt.Println(string(mpB))

	// 反序列化

	// 初始化一个结构体实例
	var d Data
	// 定义一个虚拟的字节数组
	msgByt := []byte(`{
		"Name":"oceanlee",
		"Body": "hello,世界",
		"Time": 12345
	}`)
	// 将数据传递到指针位置
	if err1 := json.Unmarshal(msgByt, &d); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(d.Name)
	}

	fmt.Println(msgByt)
}
