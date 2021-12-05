package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		x, dd, z, b int
		a, y        string
	)
	b = 0
	fmt.Print("Mau berapa kali lemparan: ")
	fmt.Scanln(&z)
	fmt.Print("Masukan angka rahasia: ")
	fmt.Scanln(&x)
	for i := 1; i <= z; i++ {
		fmt.Print("Pilih Genap atau Ganjil: ")
		fmt.Scanln(&a)
		fmt.Print("Pilih Besar atau Kecil: ")

		
		fmt.Scanln(&y)
		dice := rand.NewSource(time.Now().UnixNano())
		rd := rand.New(dice)
		dd = rd.Intn(6) + 1

		if a == "genap" && dd%2 == 0 && y == "besar" && dd > 4 {
			fmt.Println("Nilai dadu: ", dd, "Anda menang")
			b++
		} else if a == "ganjil" && dd%2 != 0 && y == "kecil" && dd <= 4 {
			fmt.Println("Nilai dadu: ", dd, "Anda menang")
			b++
		} else if a == "ganjil" && dd%2 != 0 && y == "besar" && dd > 4 {
			fmt.Println("Nilai dadu: ", dd, "Anda menang")
			b++
		} else if a == "genap" && dd%2 == 0 && y == "kecil" && dd <= 4 {
			fmt.Println("Nilai dadu: ", dd, "Anda menang")
			b++
		} else {
			fmt.Println("Nilai dadu: ", dd, "Anda kalah")
		}
	}
	/*jika menang skor akan bertambah*/
	fmt.Println("Skor Anda", b, "dari", z, "lemparan")
}
