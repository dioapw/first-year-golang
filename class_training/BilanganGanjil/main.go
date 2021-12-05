package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukan sum bilangan ganjil yang diinginkan = ")
	fmt.Scan(&x)
	a := 1
	fmt.Print("Hasil= ")
	for i := 2; i <= x; i++ {
		fmt.Print(a, ",")
		a += 2
	}
	fmt.Print(a, ".")
}
