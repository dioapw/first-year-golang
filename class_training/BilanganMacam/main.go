package main

import "fmt"

func main() {
	var (
		x, max, min, total, jum, i float64
	)
	fmt.Print("Masukan angka penentu sum recordan = ")
	fmt.Scanln(&jum)
	max = -9999999999
	min = 9999999999
	for i = 0; i <= jum; i++ {
		fmt.Print("Masukan Angka = ")
		fmt.Scanln(&x)
		if x > max {
			max = x
		} else if x < min {
			min = x
		}
		total += x
	}
	fmt.Println("Rata - rata = ", total/jum)
	fmt.Println("Nilai tertinggi = ", max)
	fmt.Println("Nilai terendah = ", min)
}
