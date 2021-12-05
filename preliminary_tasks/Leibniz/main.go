package main

import "fmt"

func main() {
	var (
		pi float32
		x  int
	)
	fmt.Print("N suku pertama: ")
	fmt.Scan(&x)
	pi = 0
	y := 1
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			pi -= (1 / float32(y))
		} else {
			pi += (1 / float32(y))
		}
		y += 2
	}
	fmt.Printf("Hasil PI: %f \n", pi*4)
}
