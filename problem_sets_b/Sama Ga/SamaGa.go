package main

import (
	"fmt"
)

func main() {
	var Kata1, Kata2 string

	fmt.Scanln(&Kata1)
	fmt.Scanln(&Kata2)
	
	fmt.Println("Sama gak nih?", Kata1 == Kata2)
	fmt.Println("Gak sama nih?", Kata1 != Kata2)
}