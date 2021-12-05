package main

import (
	"fmt"
	"math"
)

func main() {
	var n, p, i float64
	fmt.Print("Masukan nilai n untuk 2 pangkat n = ")
	fmt.Scanln(&n)
	if n >= 1 {
		for i = 1; i <= n; i++ {
			p = math.Pow(2, i)
			fmt.Println("2 ^", i, " = ", p)
		}
	} else {
		fmt.Println("2 ^ 0 = ", 1)
	}
}
