package main

import "fmt"

func main () {
	var A, B, C byte
	var Text string
	
	fmt.Scanf("%c %c %c", &A, &B, &C)
	
	Text = string(A) + string(B) + string(C)
	
	fmt.Println(Text)
}