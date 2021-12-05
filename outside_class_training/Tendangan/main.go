package main

import "fmt"

func even(num int) bool {
	return num%2 == 0
}

func odd(num int) bool {
	return num%2 == 1 && num%5 != 0
}

func five(num int) bool {
	return num%5 == 0
}

func main() {
	var (
		num    int
		degea  int = 0
		marcus int = 0
	)

	for i := 1; i <= 5; i++ {
		fmt.Print("Tendangan ke-", i, " : ")
		fmt.Scan(&num)

		genap := even(num)
		ganjil := odd(num)
		lima := five(num)

		if genap {
			fmt.Println("Tendangan terlalu ke kiri atau ke kanan")
			degea++
		} else if lima {
			fmt.Println("Tendangan terlalu ke atas")
			degea++
		} else if ganjil {
			fmt.Println("Tendangan tepat sasaran")
			marcus++
		}

	}
	if marcus > degea {
		fmt.Println("Skor akhir: ", marcus, " untuk Marcus, ", degea, " untuk De Gea")
		fmt.Println("Superb, Marcus!")
	} else {
		fmt.Println("Skor akhir: ", marcus, " untuk Marcus, ", degea, " untuk De Gea")
		fmt.Println("Well done, De Gea!")
	}
}
