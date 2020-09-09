package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	// 定义客户端请求结构体
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get("http://192.168.1.123:8000")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
	rb := resp.Body
	fmt.Printf("%T", "\n", rb)
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	f, err := os.Create("/tmp/t.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
		// 按字节写入
		f.Write(scanner.Bytes())
		// 按string写入
		f.WriteString(scanner.Text())
	}
	f.Sync()

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
