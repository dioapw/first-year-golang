package main

import "fmt"

func main() {
	var uang, bil, sisa int
	fmt.Print("Nilai uang: ")
	fmt.Scan(&uang)
	fmt.Print("Total Rp 10000,- yang dipunya: ")
	fmt.Scan(&bil)
	sisa = (uang + 9999) / 10000
	fmt.Println("Total Rp 10000,- yang diperlukan: ", sisa)
	fmt.Println("Cukup? ", bil > sisa)

}
