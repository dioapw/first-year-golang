package main

import "fmt"

func main() {
	var nilai1, nilai2, nilai3 int
	jum := 0
	i := 0
	for berhasil := false; !berhasil && jum < 70; i++ {
		fmt.Print("Masukan skor panah Anda: ")
		fmt.Scan(&nilai1, &nilai2, &nilai3)
		berhasil = nilai1 == 10 && nilai2 == 10 && nilai3 == 10
		jum = jum + nilai1 + nilai2 + nilai3
	}
	fmt.Println("Selesai")
	fmt.Println("Tercapai pada set ke-", i)
}
