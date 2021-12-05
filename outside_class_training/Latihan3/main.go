package main

import "fmt"

func main() {
	var (
		a, b, c, d, e, f, x float64
		y                   float64 = 0
	)
	fmt.Print("recordkan koefisien (a,b) : ")
	fmt.Scan(&a, &b)
	fmt.Print("recordkan koefisien (c,d) : ")
	fmt.Scan(&c, &d)
	fmt.Print("recordkan koefisien (e,f) : ")
	fmt.Scan(&e, &f)
	fmt.Print("recordkan nilai x : ")
	fmt.Scan(&x)
	y = ((a / b) * (x * x)) + ((c / d) * x) + (e / f)
	fmt.Println("Nilai y adalah : ", y)
}
