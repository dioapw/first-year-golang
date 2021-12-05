package main

import (
	"fmt"
)

func main() {
	var S string
	var N, i int

	fmt.Scanf("%s %d\n", &S, &N)

	for i = 0; i < N; i++ {
		fmt.Println(S)
	}
}
