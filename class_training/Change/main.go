package main

import "fmt"

func change(a, b int) (A, B int) {
	temp := 0
	temp = a
	A = b
	B = temp
	return
}

func main() {
	a, b := 10, 5

	A, B := change(b, a)
	fmt.Println(A, B)

	A, B = change(a, b)
	fmt.Println(B, A)
}
