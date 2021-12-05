package main

import (
	"fmt"
)

func main() {
	var Menit int

	fmt.Scanln(&Menit)
	fmt.Println("Konversi ke hours dan menit :", (Menit / 60), "hours", Menit % 60, "Menit")
}