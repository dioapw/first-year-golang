package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukan bilangan acak = ")
	fmt.Scan(&x)
	for i := 0; x != 66 && x < 100; i++ {
		fmt.Print("Masukan bilangan acak = ")
		fmt.Scan(&x)
	}
	fmt.Println("Anda memasukan angka 66 atau angka tiga digit")
	fmt.Println("{Program selesai}")
}
