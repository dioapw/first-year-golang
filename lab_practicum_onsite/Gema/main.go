package main

import "fmt"

//func hitungSkor menghitung total skor dan sum soal yang dikerjakan
func hitungSkor(num int, soal *int, skor *int) {
	if num != 301 {
		*soal++
		*skor += num
	}
}

//program menghitung kemenangan perserta dalam mengerjakan soal
func main() {
	var (
		num, soal, skor    int
		name, nametemp     string
		soaltemp, skortemp int = -1, -1
	)

	for {
		fmt.Scan(&name)
		if name == "selesai" { //program akan meminta record terus sampai diberi recordan "selesai"
			break
		}
		for i := 0; i < 8; i++ { //tiap orang mengerjakan 8 soal
			fmt.Scan(&num)
			hitungSkor(num, &soal, &skor)
		}
		switch {
		case soaltemp < soal: //pemenangnya adalah yang bisa mengerjakan sum soal terbanyak
			soaltemp = soal
			nametemp = name
			skortemp = skor
			fallthrough
		case (skortemp > skor) && (soaltemp == soal): //jika sum soal yang dikerjakan sama ,maka dilihat lebih cepat mana waktu pengerjaannya
			skortemp = skor
			nametemp = name
			soaltemp = soal
		}
		soal = 0
		skor = 0
	}
	fmt.Println(nametemp, soaltemp, skortemp)
}
