package main

import "fmt"

const Fuel int = 12

func main() {
	var Hours, Speed int
	var Spent float64
	fmt.Scanln(&Hours)
	fmt.Scanln(&Speed)
	
	Spent = float64(Hours * Speed) / float64(Fuel)
	fmt.Println(Spent)
}