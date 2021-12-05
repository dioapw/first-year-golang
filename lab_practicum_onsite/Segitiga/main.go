package main

import "fmt"

func main() {
	var sisi1, sisi2, sisi3 int
	final := " "
	fmt.Print("Sisi segitiga: ")
	fmt.Scan(&sisi1, &sisi2, &sisi3)

	if sisi1 == sisi2 && sisi1 == sisi3 {
		final = "Sama Sisi"
	} else if sisi1 == sisi3 {
		final = "Sama Kaki"
	} else if sisi1+sisi2 < sisi3 {
		final = "Bukan Segitiga"
	} else {
		final = "Sembarang"
	}

	if sisi1 < 0 || sisi2 < 0 || sisi3 < 0 {
		final = "Data Salah"
	} else if sisi3 < sisi2 || sisi3 < sisi1 {
		final = "Data Salah"
	}

	fmt.Println("Jenis Segitiga: ", final)
}
