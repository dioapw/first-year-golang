package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	const time = 5720
	age := float64(0)

	for stop := false; !stop; {
		fmt.Print("Kandungan(pct) C14 benda: ")
		fmt.Scan(&x)
		for i := 1; i != 2 && x != 100; i++ {
			age = math.Log2(100/x) * time
			fmt.Printf("age benda(tahun): %.f\n", age)
		}
		stop = x == 100
	}
	fmt.Println("PROGRAM SELESAI")
}
