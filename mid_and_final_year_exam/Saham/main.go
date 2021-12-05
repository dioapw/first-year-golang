package main

import "fmt"

func main() {
	var (
		ipo   float64
		harga float64
	)
	temp := float64(0)
	total := float64(0)
	bl, br, sd := 0, 0, 0
	fmt.Print("Masukan harga IPO: ")
	fmt.Scan(&temp)
	ipo = temp
	for i := 1; i <= 5; i++ {
		fmt.Print("Harga minggu #", i, ": ")
		fmt.Scan(&harga)
		if temp < harga {
			fmt.Println("Status BULLISH")
			bl++
		} else if temp > harga {
			fmt.Println("Status BEARISH")
			br++
		} else {
			fmt.Println("Status SIDEWAYS")
			sd++
		}
		temp = harga
		total = total + harga
	}
	fmt.Println("Terjadi", bl, "BULLISH", br, "BEARISH", sd, "SIDEWAYS")
	fmt.Printf("Harga rata-rata: %.1f\n", (total+ipo)/6)
}
