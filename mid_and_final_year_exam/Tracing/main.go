package main

import "fmt"

func main() {
	var n, i, j, k, d int
	fmt.Scanln(&n)
	i = 1
	for i < n+1 {
		j = i
		k = 1
		d = 0
		for k < j {
			if j%k == 0 {
				d++
			}
			k++
		}
		fmt.Println(j, d == 2)
		i++
	}
}
