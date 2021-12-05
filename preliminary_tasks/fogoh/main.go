package main

import (
	"fmt"
	"math"
)

//F function untuk menentukan nilai f
func F(num float64) float64 {
	sum := math.Pow(num, 2)
	return sum
}

//G function untuk menentukan nilai g
func G(num float64) float64 {
	sum := num - 2
	return sum
}

//H function untuk menentukan nilai h
func H(num float64) float64 {
	sum := num + 1
	return sum
}

//FoGoH function untuk menentukan nilai fogoh
func FoGoH(num float64) float64 {
	return F(G(H(num)))
}

//program mencari nilai fogoh
func main() {
	var num float64
	fmt.Print("Masukan nilai x = ")
	fmt.Scan(&num)

	fmt.Printf("f(%.2f) = %.2f \n", num, F(num))
	fmt.Printf("g(%.2f) = %.2f \n", num, G(num))
	fmt.Printf("h(%.2f) = %.2f \n", num, H(num))
	fmt.Printf("(fogoh) (%.2f) = %.2f \n\n", num, FoGoH(num))
}
