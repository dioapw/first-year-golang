package main

import "fmt"

func main() {
	var x, y, ls, lp, lj float64
	fmt.Print("Masukan 2 buah bilangan = ")
	fmt.Scan(&x, &y)
	lp = x * y
	ls = (x * y) / 2
	lj = x * y
	fmt.Println("Luas persegi panjang = ", lp)
	fmt.Println("Luas segitiga = ", ls)
	fmt.Println("Luas jajar genjang = ", lj)
}
