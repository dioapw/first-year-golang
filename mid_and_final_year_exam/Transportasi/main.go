package main

import "fmt"

func main() {
	var (
		time, time2 float64
		tarifwaktu1 float64
	)
	tarif := float64(0)
	jarak, jarak2 := float64(1), float64(1)
	jarakawal := float64(0)
	tarifwaktu2, tarifwaktu3 := float64(1), float64(1)
	fmt.Print("Masukan pukul berapa dan jarak pemesanannya: ")
	fmt.Scan(&time, &time2, &jarak)
	jarakawal = jarak
	if time2 >= 0 && time >= 5 && time <= 9 {
		tarifwaktu1 = 5000
		tarifwaktu2 = 4500
	} else if time2 >= 0 && time >= 16 && time <= 19 {
		tarifwaktu1 = 5000
		tarifwaktu2 = 4500
	} else if time >= 19.1 && time <= 22 || time >= 9.1 && time <= 15.9 {
		tarifwaktu3 = 4000
	} else {
		fmt.Println("Maaf,kami tidak bisa melayani pesanan Anda")
		jarak = 0
	}
	if jarak > 0 && jarak <= 10 {
		jarak = tarifwaktu1
	} else if jarak > 10 && jarak <= 20 {
		jarak = tarifwaktu2
		jarak2 = tarifwaktu3
	} else if jarak > 0 && jarak <= 20 {
		jarak2 = tarifwaktu3
	}
	tarif = jarak * jarakawal * jarak2
	fmt.Println("Harganya: ", tarif)
}
