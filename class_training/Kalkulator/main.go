package main

import "fmt"

func main() {
	var num, sum int
	fmt.Scan(&num)
	sum = num * num
	fmt.Println(sum)
	fmt.Println(num%3 == 0 && num%5 == 0 && num%10 != 0)
}
