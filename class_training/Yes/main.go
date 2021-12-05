package main

import "fmt"

//Yes is a function to change number
func Yes(x, b, c int) {
	var a int
	a = a + 5
	b = 15
	c = c + 5
	fmt.Println(x, b, c)

}

func main() {
	x, y, z := 5, 10, 15
	Yes(x, y, z)

}
