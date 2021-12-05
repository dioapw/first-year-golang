package main

import "fmt"

func main() {
	var (
		sum   int
		day   int
		day2  int
		maxOb int
	)
	increase := 0
	dayStop := 0 //penghenti counter jika tidak 2 hari berturut
	max := -999999999
	total := 0 //total hari yang berhasil
	fmt.Print("record Maksimum Hari Observasi: ")
	fmt.Scan(&maxOb)                                        // harus antara 2 - 100
	for stop := false; !stop && maxOb < 2 || maxOb > 100; { //jika tidak maka akan terus berulang sampai sesuai kondisi
		fmt.Print("record Maksimum Hari Observasi: ")
		fmt.Scan(&maxOb)
		stop = maxOb > 2 && maxOb < 100
	}
	for i := 1; i <= maxOb; i++ {
		day2 = day
		fmt.Print("Day ", i, " : ")
		fmt.Scan(&day) //jika lebih dari 1000 atau bernilai negatif maka akan diberi 3 kali kesempatan
		if day < 0 || day > 1000 {
			for j := 1; j < 3; j++ {
				fmt.Print("Day ", i, " : ")
				fmt.Scan(&day) // jika sudah 3 kali tapi nilai belum memenuhi akan otomatis diberi nilai 0
				if day > 0 && day < 1000 {
					break
				}
				if j == 2 {
					day = 0
				}
			}
		}
		sum += day
		if day > day2 && i != 1 {
			increase++ //menghitung kenaikan
		}
		if day < 10 {
			dayStop++ //jika record dibawah 10 tapi bukan negatif selama 2 hari berturut maka loop berakhir
			total = i
		} else {
			dayStop = 0
		}
		if dayStop == 2 {
			break
		}
		if day > max {
			max = day //menghitung hari terbesar
		}
		total = i
	}
	fmt.Printf("Average: %.f\n", float64(sum)/float64(total))
	fmt.Println("Max: ", max)
	fmt.Println("Number of increase: ", increase)
}
