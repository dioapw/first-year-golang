package main

import (
	"fmt"
	"math"
)

//Point digunakan untuk menyimpan nilai titik x dan y
type Point struct {
	x float64
	y float64
}

//variabel global
var (
	n         = 1000000
	points    = make([]Point, n)
	numpoints int
)

//jarak merupakan fungsi untuk menghitung jarak antar titik
func jarak(p1, p2 Point) float64 {
	sum := 0.0
	sum = math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
	return sum
}

//bacaTitik merupakan fungsi untuk menerima recordan dari user kemudian disimpan di di array struct
func bacaTitik() {
	for i := 0; i < len(points); i++ {
		fmt.Scan(&points[i].x, &points[i].y)
		if points[i].x == 0.0 && points[i].y == 0.0 {
			break
		}
		numpoints++
	}
}

//ambilTitikTerdekat merupakan fungsi untuk menentukan titik terdekat
func ambilTitikTerdekat(p1, p2 *Point) {
	var temp1 float64
	var P1, P2 Point

	p1.x, p1.y = points[0].x, points[0].y
	p2.x, p2.y = points[1].x, points[1].y

	for j := 0; j < numpoints; j++ {
		for i := 1; i < numpoints; i++ {
			P1.x, P1.y = points[j].x, points[j].y
			P2.x, P2.y = points[i].x, points[i].y
			if temp1 < jarak(P1, P2) && i != 1 {
				p1.x, p1.y = points[j].x, points[j].y
				p2.x, p2.y = points[i].x, points[i].y
				temp1 = jarak(P1, P2)
			} else {
				temp1 = jarak(P1, P2)
			}
		}
	}
}

//program menentukan titik terdekat tiap record dan menghitung jaraknya
func main() {
	var titikA, titikB Point
	bacaTitik()
	ambilTitikTerdekat(&titikA, &titikB)
	Distance := jarak(titikA, titikB)
	fmt.Printf("Titik terdekat adalah (%.1f,%.1f) dan (%.1f,%.1f) dengan jarak %.1f\n", titikA.x, titikA.y, titikB.x, titikB.y, Distance)
}
