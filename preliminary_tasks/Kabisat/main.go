package main

import "fmt"

func main() {
	var (
		tahun   int
		kabisat bool
	)
	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)
	kabisat = tahun%400 == 0 && tahun%4 == 0 || tahun%100 != 0 && tahun%4 == 0
	fmt.Printf("Apakah tahun %d kabisat? %v\n", tahun, kabisat)
}
