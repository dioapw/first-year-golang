package main

import "fmt"

func main() {
	var record int
	a := 0
	b := 1
	c := 1
	fmt.Print("record: ")
	fmt.Scan(&record)
	fmt.Println("Bilangan fibonnaci pertama: ", b)
	for i := 0; i < record; i++ {
		a = b
		b = c
		c = b + a
		fmt.Println("Bilangan fibonnaci berikutnya", b)

	}
	// contoh yang berbeda
	var recordEmployee int
	a1 := 0
	b1 := 1
	c1 := 1
	fmt.Print("record: ")
	fmt.Scan(&recordEmployee)
	fmt.Print(b1, "")
	for i := 1; i < recordEmployee; i++ {
		a1 = b1
		b1 = c1
		c1 = b1 + a1
		fmt.Print(" ", b1)

	}
}
