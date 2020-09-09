package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := "0.0.0.0:8001"
	http.HandleFunc("/", handler)
	fmt.Println("server starting: ", host)
	log.Fatal(http.ListenAndServe(host, nil))

}

type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	// 文字
	// bytes1 := []byte("hello,world!\n你好世界\n")
	// w.Write(bytes1)

	// 图片
	// wp, _ := ioutil.ReadFile("/data/linux_conf_files/wp.jpg")
	// w.Write(wp)

	// json数据
	jsonStr := Person{
		Name: "liyang",
		Age:  32,
	}

	jsonByte, _ := json.Marshal(jsonStr)
	w.Write(jsonByte)

}
