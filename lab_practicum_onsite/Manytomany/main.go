package main

import "fmt"

func main() {

	var (
		x, p, jum, jum2 int
		rahmad, rbadrun int
	)

	fmt.Print("Masukan bilangan penentu sum recordtan Ahmad : ")
	fmt.Scan(&x)
	for i := 1; i <= 3*x; i++ {
		fmt.Print("Masukan bilangan Ahmad untuk direcord : ")
		fmt.Scan(&p)
		jum = jum + p
		rahmad = jum / (3 * x)
	}
	fmt.Print("Masukan bilangan penentu sum recordtan Badrun : ")
	fmt.Scan(&x)
	for i := 1; i <= 3*x; i++ {
		fmt.Print("Masukan bilangan Badrun untuk direcord : ")
		fmt.Scan(&p)
		jum2 = jum2 + p
		rbadrun = jum2 / (3 * x)
	}
	fmt.Println("Rata-rata keseluruhan Ahmad:", rahmad, ".Yang menang? ", rahmad > rbadrun)
	fmt.Println("Rata-rata keseluruhan Badrun:", rbadrun, ".Yang menang? ", rbadrun > rahmad)
}
