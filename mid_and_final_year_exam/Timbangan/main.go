package main

import (
	"fmt"
	"math"
)

//func hitungVolume untuk menghitung volume tabung
func hitungVolume(r, t float64) float64 {
	v := 0.0
	v = math.Pi * math.Pow(r, 2) * t
	return v
}

//func hitungMassa untuk menghitung massa tabung
func hitungMassa(r, t, p float64) float64 {
	mass := 0.0
	mass = hitungVolume(r, t) * p
	return mass
}

//memproses apakah massa1 dan massa2
func display(m1, m2 float64) {
	if m1 == m2 { //jika kedua massa sama
		fmt.Println("Balance")
	} else { //jika tidak maka akan dioutputkan selisihnya
		fmt.Printf("%.2f\n", math.Abs(m1-m2))
	}
}

//program timbangan 2 benda dengan wadah berbentuk tabung
func main() {
	var (
		r               float64
		tkiri, tkanan   float64
		mjkiri, mjkanan float64
	)
	fmt.Scanln(&r)
	fmt.Scanln(&tkiri, &mjkiri)   //ada tinggi kiri dan massa jenis kiri
	fmt.Scanln(&tkanan, &mjkanan) //ada tinggi kanan dan massa jenis kanan

	masskiri := hitungMassa(r, tkiri, mjkiri)
	masskanan := hitungMassa(r, tkanan, mjkanan)
	display(masskiri, masskanan) //mengoutputkan hasil proses massa tersebut
}
