package main

import "fmt"

func main() {
	var (
		a1, a2, a3 float64
		b1, b2, b3 float64
		rahmad     float64
		rbadrun    float64
		h1         bool
		h2         bool
		h3         bool
	)
	fmt.Println("Enter 3 nilai Ahmad!")
	fmt.Scanln(&a1, &a2, &a3)
	fmt.Println("Enter 3 nilai Badrun!")
	fmt.Scanln(&b1, &b2, &b3)
	rahmad = (a1 + a2 + a3) / 3
	rbadrun = (b1 + b2 + b3) / 3
	h1 = (rahmad > rbadrun)
	h2 = (a1 > b1 && a2 > b2) || (a2 > b2 && a3 > b3) || (a1 > b1 && a3 > b3)
	h3 = (a1 < b1 && a2 < b2) || (a2 < b2 && a3 < b3) || (a1 < b1 && a3 < b3)
	fmt.Println("Nilai Rata-rata", rahmad, rbadrun)
	fmt.Println("Rata-rata Ahmad lebih baik dari Badrun ?", h1)
	fmt.Println("Apakah Ahmad pemenang absolut?", h2)
	fmt.Println("Apakah Badrun Pemenang absolut?,", h3)
}
