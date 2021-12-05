package main

import (
	"fmt"
)

func main() {
	var SukuAwal, Beda, N int

	fmt.Scanln(&N, &SukuAwal, &Beda)

	for i := 0; i < N; i++ {
		fmt.Println(SukuAwal + (Beda * i))
	}
}
