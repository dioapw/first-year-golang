package main

import "fmt"

const Phi float64 = float64(22) / float64(7)

func main() {
	var r int

	var Area float64
	
	fmt.Scanln(&r)
	
	Area = Phi * float64(r * r)
	
	fmt.Println(Area)
}