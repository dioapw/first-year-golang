package main

import (
	"fmt"
	"math"
)

//counting the distance fo the park
func distance(x1, y1, x2, y2 float64) float64 {
	sum := math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
	sum = math.Sqrt(sum)

	return sum
}

//counting if the circle points are inside or outside
func inOrout(x, y, xc, yc, r float64) bool {
	status := false
	sum := math.Pow((x-xc), 2) + math.Pow((y-yc), 2)
	if sum < math.Pow(r, 2) {
		status = true
	}
	return status
}

//program that counts distance and in or out points of a circle
func main() {
	var x1, x2, y1, y2 float64
	var xc, yc, r float64
	fmt.Print("(Titik1x, Titik1y): ")
	fmt.Scan(&x1, &y1)
	fmt.Print("(Titik2x, Titik2y): ")
	fmt.Scan(&x2, &y2)

	sum := distance(x1, y1, x2, y2)
	fmt.Printf("Jarak titik1 ke titik2 adalah: %.2f\n", sum)

	fmt.Print("(PusatX, PusatY),R: ")
	fmt.Scan(&xc, &yc, &r)
	fmt.Print("(TitikX, TitikY): ")
	fmt.Scan(&x1, &y1)

	InOrOut := inOrout(x1, y1, xc, yc, r)
	if InOrOut == true {
		fmt.Printf("Titik (%.1f, %.1f) berada di DALAM lingkaran\n", x1, y1)
	} else {
		fmt.Printf("Titik (%.1f, %.1f) berada di LUAR lingkaran\n", x1, y1)
	}
}
