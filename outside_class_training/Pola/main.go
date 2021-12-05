package main

import "fmt"

func main() {
	var N int
	fmt.Print("N : ")
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			switch {
			case i == 0 || j == 0:
				fmt.Print("*")
			case N == 2:
				fmt.Print(" ")
			case j == N-1 || i == N-1:
				fmt.Print("*")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
