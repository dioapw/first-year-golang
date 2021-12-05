package main

import "fmt"

func main() {
	var money int

	fmt.Print("Enter sum mata uang indonesia: ")
	fmt.Scanln(&money)
	for money%25 != 0 {
		fmt.Println(money, "Mata uang tidak valid.")
		fmt.Print("Enter sum mata uang indonesia: ")
		fmt.Scanln(&money)
	}
	fmt.Println(money, "Mata uang valid.")
	fmt.Println("{Program Selesai}")

	/*different section*/

	var a int
	fmt.Print("Nilai uang: ")
	fmt.Scanln(&a)
	b := 0
	for a >= 10000 {
		a -= 10000
		b++
	}
	fmt.Println("Banyaknya Rp 10.000: ", b)
}
