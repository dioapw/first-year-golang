package main

import (
	"fmt"
	"strconv"
)

//program that shows digits of each numbers
func main() {
	var num string
	var number int

	fmt.Scan(&number)
	num = strconv.Itoa(number)

	if number >= 100 && number <= 999 {
		for i, digit := range num {
			if i == 2 {
				fmt.Printf("%c \n", digit)
			} else {
				fmt.Printf("%c - ", digit)
			}
		}
	}

	// without looping
	var numbers, num1, num2, num3 int

	fmt.Scan(&numbers)

	if numbers >= 100 && numbers <= 999 {
		num1 = (numbers / 100) % 10
		num2 = (numbers / 10) % 10
		num3 = numbers % 10
		fmt.Printf("%d - %d - %d \n", num1, num2, num3)
	}
}
