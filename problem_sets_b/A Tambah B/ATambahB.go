package main

import "fmt"

func main() {
	var A, B, Sum int
	
	fmt.Scanln(&A, &B)
	
	Sum = A + B
	
	fmt.Println(Sum)
}