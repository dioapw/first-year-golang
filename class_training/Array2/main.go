package main

import "fmt"

func main() {
	A := [3][4]int{}
	k := 1
	for i := 0; i < len(A); i++ {
		for j := 0; j <= len(A); j++ {
			A[i][j] = k
			k++
		}
		fmt.Println(A[i])
	}
}
