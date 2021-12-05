package main

import "fmt"

func main() {
	var A, B, C, D, E, F float32

	fmt.Print("Masukan angka pertama: ")
	fmt.Scanln(&A)
	fmt.Print("Masukan angka kedua: ")
	fmt.Scanln(&B)

	C = A + B
	D = A - B
	E = A * B
	F = A / B

	fmt.Println("Hasil pensuman:", C)
	fmt.Println("Hasil pengurangan:", D)
	fmt.Println("Hasil perkalian:", E)
	fmt.Printf("Hasil pembagian: %.3f\n", F)

}
