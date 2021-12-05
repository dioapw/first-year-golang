package main

import "fmt"

func main() {
	var berat float64
	var p, l, t int

	sum := float64(0)
	totalP := 0
	fmt.Print("Informasi Paket: ")
	fmt.Scan(&berat, &p, &l, &t)

	totalP = p + (2 * (l + t))

	if p > 150 || l > 150 || t > 150 {
		fmt.Println("Berat Terhitung: Terlalu Panjang")
	} else if totalP > 400 {
		fmt.Println("Berat terhitung: Terlalu Besar")
	} else if berat > 50 {
		fmt.Println("Berat terhitung: Terlalu Berat")
	} else {
		sum = (float64(p) * float64(l) * float64(t)) / 6000
		if sum > berat {
			fmt.Printf("Berat terhitung: %.3f\n", sum)
		} else {
			fmt.Printf("Berat terhitung: %.1f\n", float64(berat))
		}
	}
}
