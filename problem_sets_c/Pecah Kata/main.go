package main

import "fmt"

func main() {
	var Text string

	fmt.Scanln(&Text)

	for i, word := range Text {
		fmt.Printf("Karakter ke-%d adalah %c\n", i+1, word)
	}
}
