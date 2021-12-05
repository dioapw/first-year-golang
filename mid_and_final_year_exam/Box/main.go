package main

import "fmt"

//goods transportation program
func main() {
	var weight1, weight2 int
	var sum, sumW1, sumW2 int
	// program will asked you to give an record until the sum > 150
	for stop := false; !stop; {
		fmt.Print("Masukan beban kedua box: ")
		fmt.Scan(&weight1, &weight2)
		sumW1 += weight1
		sumW2 += weight2
		sum = sum + weight1 + weight2
		stop = sum > 150
	}
	//weight comparison
	if sumW1 > sumW2 {
		fmt.Println("Beban berat kekiri")
	} else if sumW2 > sumW1 {
		fmt.Println("Beban berat kekanan")
	} else {
		fmt.Println("Seimbang")
	}
}
