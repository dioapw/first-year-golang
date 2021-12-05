package main

import "fmt"

func main() {
	var i, j float64
	fmt.Print("Masukan sum penumpang = ")
	fmt.Scan(&j)
	i = 0
	for j/7 > 0 {
		j -= 7
		i++
	}
	fmt.Println("sum minibus yang harus digunakan = ", i)
}
