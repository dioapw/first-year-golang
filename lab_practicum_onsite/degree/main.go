package main

import "fmt"

func main() {
	var f, r, k float64
	var degree int

	fmt.Print("Enter nilai suhu celcius: ")
	fmt.Scan(&degree)

	f = (float64(9) * float64(degree) / float64(5)) + 32
	r = (float64(4) * float64(degree) / float64(5))
	k = float64(degree) + float64(273.15)

	fmt.Println("Nilai konversi suhu : ")
	fmt.Printf("Dalam Reamur = %.2f\n", r)
	fmt.Printf("Dalam Fahrenheit =  %.2f\n", f)
	fmt.Printf("Dalam Kelvin = %.3f\n", k)
}
