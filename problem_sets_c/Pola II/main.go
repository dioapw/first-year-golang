package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		for j := 0; j+i < N-1; j++ {
			fmt.Print(" ")
		}

		for j := 0; j <= (i * 2); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
