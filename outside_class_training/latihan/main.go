package main

import (
	"fmt"
	"math"
)

func main() {
	// About a fibonacci ID
	maxF := 100
	f0 := 0
	f1 := 1
	f2 := 1
	fmt.Println("Bilangan fibonacci pertama: ", f1)
	for selesai := false; !selesai; {
		f0 = f1
		f1 = f2
		f2 = f1 + f0
		fmt.Println("Bilangan fibonacci berikutnya: ", f1)
		selesai = f2 > maxF

	}
	//math packages
	fmt.Println(math.Floor(16.4))
	fmt.Println(math.Ceil(16.4))
	fmt.Println(math.Sqrt(16))
}
