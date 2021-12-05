package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a, x, dd, dn int
	fmt.Print("Masukan angka rahasia: ")
	fmt.Scanln(&x)
	fmt.Print("Masukan angka tebakan Anda: ")
	fmt.Scanln(&a)

	dice := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(dice)
	dd = rd.Intn(6) + 1
	dn = rd.Intn(6) + 1
	fmt.Println("Danang: ", dd)
	fmt.Println("Nilai dadu: ", dn)
	if a == dd && dn == dd {
		fmt.Println("Seri")
	} else if dn == dd {
		fmt.Println("Danang Menang")
	} else if dn == a {
		fmt.Println("Anda Menang")
	} else {
		fmt.Println("tidak ada pemenang")
	}

}
