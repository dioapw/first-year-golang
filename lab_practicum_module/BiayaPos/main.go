package main

import "fmt"

func main() {
	var gr, kg, berat int
	var hargakg, hargagr, harga int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	gr = berat % 1000
	kg = berat / 1000
	fmt.Println("Detail berat: ", kg, "kg + ", gr, "gr")
	hargakg = 10000 * kg
	if gr >= 500 {
		hargagr = (5 * gr)
		harga = hargakg + hargagr
	} else if gr <= 500 {
		hargagr = (15 * gr)
		harga = hargakg + hargagr
	}
	fmt.Println("Detail biaya: Rp ", hargakg, " + Rp ", hargagr)
	if kg > 10 {
		harga = hargakg
	}
	fmt.Println("Total biaya: Rp ", harga)
}
