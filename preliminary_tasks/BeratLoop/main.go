package main

import "fmt"

func main() {
	var a, b float64 = 0.0, 0.0
	for a < 9 && b < 9 {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&a, &b)
	}
	fmt.Println("Proses selesai.")
}
