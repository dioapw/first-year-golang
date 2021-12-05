package main

import "fmt"

func main() {
	var nilai, min int

	min = 999999999
	i := 1
	totalM := 0

	fmt.Print("Mahasiswa ", i, " : ")
	fmt.Scan(&nilai)

	for stop := false; !stop; {
		totalM++
		if nilai < 0 || nilai > 700 {
			for no := false; !no; {
				fmt.Print("Mahasiswa ", i, " : ")
				fmt.Scan(&nilai)
				no = nilai > 0 && nilai < 700
			}
		}
		if nilai < min {
			min = nilai
		}
		i++
		fmt.Print("Mahasiswa ", i, " : ")
		fmt.Scan(&nilai)
		stop = (nilai == -999)
	}

	fmt.Println("SELESAI")
	fmt.Println("sum MAHASISWA: ", totalM)
	fmt.Println("MIN: ", min)
}
