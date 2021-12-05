package main

import "fmt"

func main() {
	var (
		pi, pi2, jp float64
		i           int
	)
	pi = 0
	pi2 = 0
	y := 1
	i = 1
	for berhenti := false; !berhenti; {

		pi2 = pi

		if i == 1 {
			pi = pi + 1
		} else {
			y += 2
			if i%2 == 0 {
				pi -= (1 / float64(y))
			} else {
				pi += (1 / float64(y))
			}
		}
		jp = pi*4 - pi2*4

		if jp < 0 {
			jp = jp * -1
		}
		berhenti = jp < 0.00001
		i++
	}
	fmt.Printf("Hasil PI: %.11f \n", pi2*4)
	fmt.Printf("Hasil PI: %.11f \n", pi*4)
	fmt.Println("Pada i ke : ", i)
}
