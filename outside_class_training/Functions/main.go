package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	sum := getSum(5, 6)
	name := greeting("Dio")
	fmt.Println("Bonjour atau", name)
	fmt.Println("5 + 6 = ", sum)
}
