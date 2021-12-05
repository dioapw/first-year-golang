package main

import "fmt"

//example :
//1,-2,3,-4 = -2
func main() {
	var n, sum, sum2 int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum2 = i * -1
			fmt.Print(",", sum2)
			sum += sum2
		} else {
			if i != 1 {
				fmt.Print(",", i)
			} else {
				fmt.Print(i)
			}
			sum += i
		}
	}
	fmt.Println(" =", sum)
}
