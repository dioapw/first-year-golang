package main

import "fmt"

func main() {
	var a, b, c, d, e byte
	var b1 = []byte{104, 97, 108, 108, 111, 111}
	var s = string(b1)

	fmt.Printf("%s \n", s)
	//different section
	fmt.Println("Enter Angka untuk diubah ke ASCII !")
	fmt.Scanln(&a, &b, &c, &d, &e)
	fmt.Printf("%c%c%c%c%c", a, b, c, d, e)
	fmt.Println()
	fmt.Printf("%c%c%c", a+1, b+1, c+1)

}
