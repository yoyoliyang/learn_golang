package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午好")
	case t.Hour() < 17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}
}
