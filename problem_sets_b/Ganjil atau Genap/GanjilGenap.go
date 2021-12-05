package main

import (
	"fmt"
)

func main() {
	var N int
	
	fmt.Scanln(&N)
	
	fmt.Println("Ganjil?", N % 2 == 1)
	fmt.Println("Genap?", N % 2 == 0)
}