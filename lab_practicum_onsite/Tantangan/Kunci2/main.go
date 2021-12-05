package main

import "fmt"

func main() {
	var x, nilai, nilai2 int
	jum := 0
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		fmt.Scan(&nilai, &nilai2)
		jum = nilai + nilai2
		if jum%2 == 0 {
			fmt.Println(nilai2)
		} else {
			fmt.Println(nilai)
		}

	}

}
