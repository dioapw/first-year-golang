package main

import "fmt"

func main() {
	var b int
	fmt.Print("Masukan angka = ")
	fmt.Scan(&b)
	i := 0
	for b != 999 {
		fmt.Print("Masukan angka = ")
		fmt.Scan(&b)
		i += b
	}
	fmt.Print("Hasil akhir = ", i)
}
