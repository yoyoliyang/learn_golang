package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	const body = "hello,world!\n你好，世界！\n第三行\n第四行"

	scanner := bufio.NewScanner(strings.NewReader(body))

	for i := 0; scanner.Scan() && i < 3; i++ {
		fmt.Println(i, scanner.Text())
	}
}
