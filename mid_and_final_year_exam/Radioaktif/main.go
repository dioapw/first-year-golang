package main

import "fmt"

func main() {
	var x, x2 float64
	aktif := 0
	pasif := 0
	i := 1

	fmt.Println("Observasi unsur X")
	for stop := false; !stop; {
		fmt.Printf("Temperatur # %d : ", i)
		fmt.Scan(&x)
		stop = (x == x2)
		if i == 1 || x > x2 {
			fmt.Println("Status Aktif")
			aktif++
		} else if x < x2 {
			fmt.Println("Status Pasif")
			pasif++
		}
		x2 = x
		i++
	}

	fmt.Printf("Teramati %d aktif dan %d  pasif", aktif, pasif)

}
