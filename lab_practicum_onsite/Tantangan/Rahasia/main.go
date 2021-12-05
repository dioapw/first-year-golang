package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nilai, nilai2, angka int
	jum := 0
	acak := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(acak)
	fmt.Scan(&angka)
	fmt.Println(angka)
	for i := 0; i < angka; i++ {
		fmt.Scan(&nilai)
		nilai2 = rd.Intn(149) + 1
		jum = nilai + nilai2
		if jum%2 == 0 {
			fmt.Println(nilai, nilai2)
		} else {
			fmt.Println(nilai2, nilai)
		}
	}
}
