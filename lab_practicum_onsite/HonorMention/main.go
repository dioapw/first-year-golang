package main

import "fmt"

func main() {
	var (
		b, z, a float64
		x, c, l int
	)
	fmt.Print("Ada berapa banyak tim: ")
	fmt.Scanln(&x)
	fmt.Print("Masukan berapa minimum rata-rata honorable mention: ")
	fmt.Scanln(&a)
	z = 0
	c = 0
	l = 0
	for i := 1; i <= x; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print("tim ke ", i, " Masukan angka predikat tim mu: ")
			fmt.Scanln(&b)
			z += b
		}
		if z/3 > a {
			c++
		} else {
			l++
		}
		z = 0
	}
	/*yang melebihi minimum rata-rata mendapatkan honorable mention*/
	fmt.Println("Yang mendapatkan Honorable Mention sebanyak", c, "tim")
	fmt.Println("Yang tidak mendapatkan Honorable Mention sebanyak", l, "tim")

}
