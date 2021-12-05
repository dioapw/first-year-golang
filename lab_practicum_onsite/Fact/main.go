package main

import "fmt"

//func fact untuk menghitung faktorial
func fact(num int) int {
	sum := 1
	num2 := num
	for i := 0; i < num2; i++ {
		sum *= num
		num--
	}
	return sum
}

//program untuk menghitung faktorial
func main() {
	var num, first, sumNum int
	i := 1
	sum := ""
	for {
		fmt.Scanln(&num)
		sum = fmt.Sprintf("%s%d\n", sum, fact(num)) //menyimpan hasil faktorial
		if i == 1 {
			first = num
		} else { //program berhenti jika selisih angka selanjutnya lebih besar dari angka pertama
			sumNum += num
			if sumNum >= first {
				break
			}
		}
		i++

	}
	fmt.Println()
	fmt.Println(sum)
}
