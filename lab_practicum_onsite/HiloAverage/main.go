package main

import "fmt"

func main() {
	var (
		b, z, a, y, v               float64
		x                           int
		name, namemenang, namekalah string
	)
	fmt.Print("Ada berapa banyak tim: ")
	fmt.Scanln(&x)
	z = 0
	a = 0
	v = -9999999999
	y = 9999999999
	for i := 1; i <= x; i++ {
		fmt.Print("Masukan name tim: ")
		fmt.Scan(&name)
		fmt.Print("Masukan angka predikat tim mu: ")
		for j := 1; j <= 3; j++ {
			fmt.Scan(&b)
			z += b
		}
		a = z / 3
		if a > v {
			v = a
			namemenang = name
		} else if a < y {
			y = a
			namekalah = name
		}
		z = 0
		a = 0
		if i == x {
			fmt.Scan(&name, &b)
		}
	}
	/* tim dengan rerata tertinggi lah yang menang dan sebaliknya*/
	fmt.Println("Tim rerata faktor tertinggi adalah tim", namemenang, "dengan nilai", v)
	fmt.Println("Tim rerata faktor terendah adalah tim", namekalah, "dengan nilai", y)
}
