package main

import (
	"fmt"
	"math"
)

func isConsecutive(num int) bool {
	digit1, digit2, selisih, i := 0, 0, float64(0), 1
	flag := true
	for num > 0 {
		digit1 = num % 10
		selisih = math.Abs(float64(digit1) - float64(digit2))
		if selisih != 1 && i != 1 {
			flag = false
		}
		num = num / 10
		digit2 = digit1
		i++
	}
	return flag

}

func main() {
	var num int = 1

	for num > 0 {
		fmt.Print("record: ")
		fmt.Scan(&num)
		if num > 9 {
			if isConsecutive(num) {
				fmt.Println("konsekuetif")
			} else {
				fmt.Println("Bukan konsekuetif")
			}
		}
	}
}
