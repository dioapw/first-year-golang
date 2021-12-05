package main

import "fmt"

func change(original *int, value int) {
	*original = value
}

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T \n", b)

	// use * to read val from addres

	fmt.Println(*b)
	fmt.Println(&a)

	// change val with pointer

	*b = 10
	fmt.Println(a)

	// parameter

	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10

}
