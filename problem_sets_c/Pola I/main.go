package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanln(&N)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
