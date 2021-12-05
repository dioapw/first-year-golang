package main

import "fmt"

func main() {
	var k, temp, fktotal float64
	fktemp := float64(1)
	akar2 := float64(1)
	fk := float64(1)
	fmt.Print("Nilai k: ")
	fmt.Scan(&k)
	for y := float64(0); y <= k; y++ {
		for i := 1; i <= 2; i++ {
			temp = (4 * y) + 2
			fk = fk * temp
		}
		for j := 1; j < 2; j++ {
			for z := 1; z < 2; z++ {
				temp = (4 * y) + 1
				fktemp *= temp
			}
			temp = (4 * y) + 3
			fktemp *= temp
		}
		fktotal = fk / fktemp
		akar2 *= fktotal
		fk = 1
		fktemp = 1
	}
	fmt.Printf("Nilai akar 2: %.10f\n", akar2)

}
