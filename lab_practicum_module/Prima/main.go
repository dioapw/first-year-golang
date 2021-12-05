package main

import "fmt"

func getprima(j int) bool { // menentukan apakah bilangan prima atau bukan
	status := false
	status = j == 2
	return status
}

func main() {
	var (
		nilai int
		z     int = 1 //counter posisi prima
		j     int = 0 //counter penentu bilangan prima
		prima int = 0
		sum   int = 0 //sum bilangan prima
		sumz  int = 0 //sum posisi prima
	)

	for stop := false; !stop; { //akan mengrecord terus sampai terdapat 3 bilangan prima
		fmt.Scanln(&nilai)
		for i := 1; i <= nilai; i++ {
			if nilai%i == 0 {
				j++
			}
		}
		if getprima(j) == true {
			prima++
			sum += nilai
			sumz += z
		}
		stop = prima == 3
		j = 0
		z++
	}
	fmt.Println("sum bilangan prima adalah ", sum)
	fmt.Println("sum posisi prima ditemukan adalah", sumz)
}
