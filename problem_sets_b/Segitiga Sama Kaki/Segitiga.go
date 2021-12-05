package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)
	
	fmt.Println("Is this Segitiga Sama sisi?", (A == B) && (B == C))
}