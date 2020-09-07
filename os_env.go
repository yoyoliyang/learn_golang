package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	os.Setenv("FOO", "bar")
	fmt.Println(os.Getenv("FOO"))

	// fmt.Println(os.Environ())

	for _, e := range os.Environ() {
		// pair := strings.SplitN(e, "=", 2)
		// fmt.Println(pair[0], pair[1])
		fmt.Println(e)
	}

}
