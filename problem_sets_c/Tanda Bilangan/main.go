package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanln(&N)

	if N > 0 {
		fmt.Println("Positif")
	} else if N < 0 {
		fmt.Println("Negatif")
	} else {
		fmt.Println("Nol")
	}
}
