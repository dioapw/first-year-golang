package main

import "fmt"

func main() {
	//factor ID
	var i int
	fmt.Print("Enter angka untuk difaktorialkan :")
	fmt.Scanln(&i)
	fmt.Print(i, "!=")
	jum := 1
	for i >= 2 {
		fmt.Print(i, "x")
		jum = jum * i
		i--
	}
	fmt.Print("1=", jum)
}
