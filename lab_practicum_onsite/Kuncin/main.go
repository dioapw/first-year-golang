package main

import "fmt"

func main() {
	var x, nilai, nilai2 int
	fmt.Print("Masukan angka rahasia: ")
	fmt.Scan(&x)
	if x < 0 {
		fmt.Println("{PROGRAM SELESAI}")
	}
	for i := 0; i < x; i++ {
		fmt.Print("Masukan 2 bilangan acak: ")
		fmt.Scan(&nilai, &nilai2)
		if nilai%x == 0 {
			fmt.Println(nilai)
		} else if nilai2%x == 0 {
			fmt.Println(nilai2)
		} else {
			fmt.Println("Bukan bilangan yang diinginkan")
		}
	}
}
