package main

import "fmt"

func main() {
	var gol int
	sum, sumP, all := 0, 0, 0
	gol1, gol2, gol3 := 0, 0, 0
	//program menghitung pendapatan gerbang tol
	for i := 1; i <= 3; i++ {
		fmt.Print("Baris ", i, " : ")
		sumP = 0
		for gagal := false; !gagal; { //meEnter record terus sampai kondisi per barisnya
			sum = 0
			fmt.Scan(&gol)
			if gol == 1 {
				sum = 1000
				gol1++
			} else if gol == 2 {
				sum = 2000
				gol2++
			} else if gol == 3 {
				sum = 3000
				gol3++
			}
			sumP += sum
			gagal = (gol == -1)
		}
		all += sumP
		fmt.Println("Pendapatan baris ke ", i, " : ", sumP)
	} //total ada 3 baris
	fmt.Println()
	fmt.Println("Pendapatan total: ", all)
	fmt.Println("Gol 1 ada: ", gol1)
	fmt.Println("Gol 2 ada: ", gol2)
	fmt.Println("Gol 3 ada: ", gol3)
}
