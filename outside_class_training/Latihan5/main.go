package main

import "fmt"

//small to large program
func main() {
	var num, num2 float64
	right := 0
	j := 1
	for start := false; !start; { //program will ask an record of integer 3 times
		fmt.Printf("Perulangan ke %d : ", j)
		for i := 1; i <= 3; i++ {
			fmt.Scan(&num)
			if num2 < num && i != 1 {
				right++
			} else {
				right = 0
			}
			num2 = num
		}
		start = right == 2 //the program will stop running after 3 integer records are given sequentially from small to large
		j++
	}
	fmt.Println("Selesai")
}
