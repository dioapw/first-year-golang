package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		x, dd int
		a, y  string
	)

	fmt.Print("Masukan angka rahasia: ")
	fmt.Scanln(&x)
	fmt.Print("Pilih Genap atau Ganjil: ")
	fmt.Scanln(&a)
	fmt.Print("Pilih Besar atau Kecil: ")
	fmt.Scanln(&y)

	dice := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(dice)
	dd = rd.Intn(6) + 1
	switch {
	case a == "genap" && dd%2 == 0 && y == "besar" && dd > 4:
		fmt.Println("Nilai dadu: ", dd, "Anda menang")
	case a == "ganjil" && dd%2 != 0 && y == "kecil" && dd <= 4:
		fmt.Println("Nilai dadu: ", dd, "Anda menang")
	case a == "ganjil" && dd%2 != 0 && y == "besar" && dd > 4:
		fmt.Println("Nilai dadu: ", dd, "Anda menang")
	case a == "genap" && dd%2 == 0 && y == "kecil" && dd <= 4:
		fmt.Println("Nilai dadu: ", dd, "Anda menang")
	default:
		fmt.Println("Nilai dadu: ", dd, "Anda kalah")
	}
}
