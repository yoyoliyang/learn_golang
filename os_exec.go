package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	lsFileList := exec.Command("ls")
	// fmt.Println(lsFileList)

	lsOutput, err := lsFileList.Output()
	if err != nil {
		panic(err)
	}

}
