package main

import "fmt"

func main() {
	var x int
	var prima bool
	j := 0
	prima = true
	fmt.Print("Masukan bilangan untuk difaktorkan: ")
	fmt.Scanln(&x)
	fmt.Print("Hasil: ")
	for i := 1; x >= i; i++ {
		if x == i {
			fmt.Print(i)
			j++
		} else if x%i == 0 {
			fmt.Print(i, ",")
			j++
		}
	}
	if j == 2 {
		fmt.Println()
		fmt.Println("Prima: ", prima)
	} else {
		prima = false
		fmt.Println()
		fmt.Println("Prima: ", prima)
	}
}
