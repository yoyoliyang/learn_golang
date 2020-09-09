package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("hello,world!\n你好世界")

	p := make([]byte, 5)

	for {
		// io.Read返回的是字节数n
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF: ", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, p[:n])

	}

	fmt.Println("________________________")

	bytesBf1 := new(bytes.Buffer)
	bytes1 := []byte("hello")
	bytesBf1.Write(bytes1)
	fmt.Println(bytesBf1.String())

	fmt.Fprintf(bytesBf1, "%T", bytes1)
}
