package main

import (
	"fmt"
)

func main() {
	temp(5)
	// fmt.Println(s)
}

func temp(n int) {
	start := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}
		for j := 1; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}
}
