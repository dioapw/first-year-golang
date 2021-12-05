package main

import "fmt"

func main() {
	var nilai, nilai1 int
	faktor := 0
	sama := 0

	fmt.Print("record user: ")
	fmt.Scan(&nilai)
	nilai1 = nilai
	for i := 1; nilai != 9999 && nilai1 != 9999; i++ {
		if i == 1 {
			for j := 1; j <= nilai1; j++ {
				if nilai1%j == 0 {
					faktor++
				}
			}
		}
		if nilai1 == nilai {
			sama++
		}
		fmt.Scan(&nilai)
	}
	if nilai1 == 9999 {
		fmt.Println("{PROGRAM SELESAI}")
	} else {
		fmt.Println(nilai1, sama, faktor == 3)
		fmt.Println("{PROGRAM SELESAI}")
	}
}
