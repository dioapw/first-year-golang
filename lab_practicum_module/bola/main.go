package main

import "fmt"

func main() {
	var r float64
	const phi = 3.1415926535
	var luasbola float64
	var volumebola float64

	fmt.Print("Jejari: ")
	fmt.Scan(&r)
	luasbola = (4 * phi * r * r)
	volumebola = (4 * phi * r * r * r) / 3
	fmt.Printf("luas bola dengan = %.4f \n", luasbola)
	fmt.Printf("volume bola = %.4f \n", volumebola)
}
