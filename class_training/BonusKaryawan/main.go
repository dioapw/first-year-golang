package main

import "fmt"

func main() {
	var c, e int
	var f string
	fmt.Print("Masukan waktu kerja per tahun dan agenya = ")
	fmt.Scanln(&c, &e)
	if c >= 5 && e >= 50 {
		f = "1 juta"
	} else if c >= 5 && e < 50 {
		f = "800 ribu"
	} else {
		f = "500 ribu"
	}
	fmt.Println("Bonus Anda = ", f)
}
