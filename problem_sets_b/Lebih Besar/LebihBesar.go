package main

import (
	"fmt"
)

func main() {
	var Bil1, Bil2 int

	fmt.Scanln(&Bil1)
	fmt.Scanln(&Bil2)
	
	fmt.Println("Apakah bilangan pertama lebih besar dari bilangan ke dua?", Bil1 > Bil2)
	fmt.Println("Apakah sebaliknya?", Bil1 <= Bil2)
}