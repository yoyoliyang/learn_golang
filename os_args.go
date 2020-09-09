package main

import (
	"fmt"
	"os"
)

func main() {
	for index, item := range os.Args {
		fmt.Println(index, item)
	}
}
