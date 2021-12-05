package main

import "fmt"

func main() {
	var Text1, Text2, Hasil string
	
	fmt.Scanln(&Text1)
	fmt.Scanln(&Text2)
	
	Hasil = Text1 + string(' ') + Text2
	
	fmt.Println(Hasil)
}