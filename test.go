package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(<-time.After(1 * time.Second))
}
