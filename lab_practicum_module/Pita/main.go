package main

import "fmt"

func main() {
	var n, a int
	var bunga, z string
	fmt.Println("Ada berapa bunga yang kau miliki? ")
	fmt.Print("N: ")
	fmt.Scan(&n)
	fmt.Print("Bunga ke- 1: ")
	fmt.Scan(&bunga)
	z = ""
	a = 0
	for l := 1; l <= n && bunga != "selesai"; l++ {
		z = fmt.Sprintf("%s%s - ", z, bunga)
		a++
		l = n
	}
	for i := 2; i <= n && bunga != "selesai"; {
		fmt.Print("Bunga ke- ", i, ": ")
		fmt.Scan(&bunga)
		for j := 1; j <= n && bunga != "selesai"; j++ {
			z = fmt.Sprintf("%s%s - ", z, bunga)
			j = n
			a++
		}
		i++
	}
	fmt.Println("Pita: ", z)
	fmt.Println("Bunga: ", a)
}
