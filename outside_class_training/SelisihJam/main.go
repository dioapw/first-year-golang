package main

import "fmt"

func main() {
	var hours, menit int
	hoursM := 0
	menitM := 0
	lamaM := 0
	lamaJ := 0
	lamaJS := 0

	fmt.Print("Mulai  (hours,menit) : ")
	fmt.Scan(&hours, &menit)

	hoursM = hours
	menitM = menit

	fmt.Print("Selesai(hours,menit) : ")
	fmt.Scan(&hours, &menit)

	lamaM = (hoursM-hours)*60 + (menitM - menit)

	if lamaM < 0 {
		lamaM *= -1
	}

	lamaJ = lamaM / 60
	lamaJS = lamaM % 60

	fmt.Printf("Lama (menit)       : %d menit\n", lamaM)
	fmt.Printf("Lama (hours,menit)   : %d hours %d menit\n", lamaJ, lamaJS)

}
