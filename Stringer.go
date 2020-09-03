package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}
func main() {
	fmt.Println("vim-go")
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"dns":      {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Println(name, ip)
	}

}
