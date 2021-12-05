package main

import "fmt"

func main() {
	var n, p int
	fmt.Print("Masukan sum soal yang ingin dikerjakan = ")
	fmt.Scan(&n)
	p = n
	for i := 1; i < n; i++ {
		p--
		fmt.Println("Ini soal ke-", i, "yang saya kerjakan,masih tersisa", p, "soal lagi")
	}
	fmt.Println("Ini soal terakhir yang saya kerjakan cukup sudah!")
}
