package main

import (
	"fmt"
	"math"
	"strings"
)

func calculate(d float64) (float64, float64) {
	// hitung luas
	area := math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	circumference := math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func calculates(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func main() {
	diameter := 15.0
	area, circumference := calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	//<---------------------------------------------------------->
	avg := calculates(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	msg := fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	//var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	//var avg = calculate(numbers...)
	//var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	//fmt.Println(msg)

	hobbies := []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)

	//yourHobbies("wick", "sleeping", "eating")
}
