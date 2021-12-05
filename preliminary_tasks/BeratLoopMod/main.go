package main

import "fmt"

func main() {
	var (
		a, b float64 = 0.0, 0.0
		c    bool
	)
	fmt.Print("Masukan berat belanjaan di kedua kantong: ")
	fmt.Scanln(&a, &b)

	for a+b <= 150 && a > -1 && b > -1 {
		fmt.Println("Sepeda motor pak Andi akan oleng: ", c)
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&a, &b)
		c = (b-a >= 9 || a-b >= 9)
	}
	fmt.Println("Program selesai.")
}
