package main

import "fmt"

func main() {
	var k, fk float64
	temp := float64(0)
	fk = 1
	fktemp := float64(1)
	fktotal := float64(0)
	fmt.Print("Nilai k: ")
	fmt.Scan(&k)
	for i := 1; i <= 2; i++ {
		temp = (4 * k) + 2
		fk = fk * temp
	}
	for j := 1; j < 2; j++ {
		for z := 1; z < 2; z++ {
			temp = (4 * k) + 1
			fktemp = fktemp * temp
		}
		temp = (4 * k) + 3
		fktemp = fktemp * temp
	}
	fktotal = fk / fktemp
	fmt.Printf("Nilai f(K): %.10f\n", fktotal)
}
