package main

import "fmt"

func main() {
	var N, Sum int

	fmt.Scanln(&N)

	Sum = 0
	for i := 0; i < N; i++ {
		Sum += (i + 1)
		fmt.Println(i + 1)
	}
	fmt.Println(Sum)
}
