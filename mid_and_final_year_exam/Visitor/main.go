package main

import "fmt"

//program menghitung sum penurunan pengunjung
func main() {
	var (
		day      int //pengunjung
		day2     int //counter penurunan pengunjung
		decrease int = 0
		over     int = 0 // counter kelebihan pengunjung
	)
	i := 1
	for j := 1; 1 > 0; j++ { //program akan mengrecord ulang terus sampai kondisi tercapai
		fmt.Print("Hari ", i, " : ")
		fmt.Scan(&day)
		if day < 0 { // jika ada record negatif maka diabaikan dan diminta mengrecord ulang
			for stop := false; !stop; {
				fmt.Print("Hari ", i, " : ")
				fmt.Scan(&day)
				stop = day > 0
			}
		}
		if day < day2 && i != 1 { //penurunan pengunjung
			decrease++
		}
		if day > 100 { //jika pengunjung lebih dari 100 2 hari berturut program berhenti
			over++
		} else {
			over = 0
		}
		if over == 2 {
			break
		}
		day2 = day
		i++
	}
	fmt.Println(decrease)
	fmt.Println(i) //pada hari ke berapa berhenti
}
