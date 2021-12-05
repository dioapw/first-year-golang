package main

import "fmt"

func maxmin(A [5]int) (max, min int) {
	max = A[0]
	min = A[0]
	for i := 2; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		}
		if min > A[i] {
			min = A[i]
		}
	}
	return
}

func main() {
	//number 1
	A := [20]int{}
	for i := 0; i < len(A); i++ {
		A[i] = i
	}
	fmt.Println(A)

	//number 2
	jum := 0
	B := [5]int{}
	for i := 0; i < len(B); i++ {
		fmt.Scan(&B[i])
		jum += B[i]
	}
	Max, Min := maxmin(B)
	fmt.Printf("Max: %d ,Min: %d\n", Max, Min)
	fmt.Println(jum)

	//number 3
	C, D := [5]int{}, [5]int{}
	for i := 0; i < len(C); i++ {
		fmt.Scan(&C[i])
	}
	fmt.Println(C)
	for i := 0; i < len(C); i++ {
		D[i] = C[i]
	}
	fmt.Println(D)

	//number 4
	fib := [21]int{}
	fib[0] = 1
	fib[1] = 1
	fib[2] = 2
	for i := 3; i < len(fib); i++ { //20 fibonnaci number
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println(fib)
}
