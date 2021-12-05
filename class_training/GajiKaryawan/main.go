package main

import "fmt"

func main() {
	var (
		ghours, ghours1, lama, slama, sisa, lembur, gtetap, total float64
		name, gol                                             string
	)

	fmt.Print("Masukan class Anda dan name Anda = ")
	fmt.Scanln(&gol, &name)
	switch gol {
	case "1":
		gtetap = 500000
		ghours = 5000
	case "2":
		gtetap = 300000
		ghours = 3000
	case "3":
		gtetap = 250000
		ghours = 2000
	case "4":
		gtetap = 100000
		ghours = 1500
	case "5":
		gtetap = 50000
		ghours = 1000
	default:
		fmt.Println("Tidak termasuk dalam class")
	}
	fmt.Print("Masukan lama waktu bekerja = ")
	fmt.Scanln(&lama)
	if lama > 150 {
		slama = lama - 150
		lembur = 1.5 * ghours * slama
		ghours1 = lama - slama
		sisa = ghours1 * ghours
		total = gtetap + sisa + lembur
	} else {
		lembur = 0
		sisa = ghours * lama
		total = gtetap + sisa + lembur
	}
	fmt.Println("name : ", name)
	fmt.Printf("Waktu: %.f hours \n", lama)
	fmt.Printf("Total Keluaran: Rp %.0f\n", total)
}
