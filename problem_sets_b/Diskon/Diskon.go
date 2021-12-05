package main

import (
	"fmt"
)

func main() {
	var Harga, Diskon int
	var Bayar float64
	
	fmt.Scanln(&Harga)
	fmt.Scanln(&Diskon)
	
	Bayar = (float64(100 - Diskon) / 100) * float64(Harga)
	fmt.Println(Bayar)
}