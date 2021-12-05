package main

import "fmt"

func main() {
	var (
		b             float64
		x, c, l, d, m int
	)
	fmt.Print("Ada berapa banyak tim: ")
	fmt.Scanln(&x)
	c = 0
	l = 0
	d = 0
	m = 0
	for i := 1; i <= x; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Scan(&b)
			if b == 100 {
				c++
			} else {
				l++
			}
		}
		if c == 3 {
			d++
		} else {
			m++
		}
		c = 0
		l = 0
		if i == x {
			fmt.Scan(&b)
		}
	}
	/*predikat luar biasa adalah yang mendapat nilai 100 semua*/
	fmt.Println("Yang mendapatkan predikat luar biasa sebanyak", d, "tim dari", x, "tim")
	fmt.Println("Yang tidak mendapatkan tidak mendapatkan predikat luar biasa sebanyak", m, "tim dari", x, "tim")

}
