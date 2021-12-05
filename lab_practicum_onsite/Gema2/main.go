package main

import "fmt"

//func hitungSkor menghitung total skor dan sum soal yang dikerjakan
func hitungSkor(num int, soal *int, skor *int) {
	if num != 301 {
		*soal++
		*skor += num
	}
}

//func disq menentukan apakah total skor lebih dari 5 hours
func disq(Skor int) bool {
	status := false
	if Skor > 300 {
		status = true
	}
	return status
}

//func valid menentukan apakah recordan negatif atau lebih dari 5 hours 1 menit
func valid(num int) bool {
	status := false
	if num < 0 || num > 301 {
		status = true
	}
	return status
}

//program menghitung kemenangan perserta dalam mengerjakan soal
func main() {
	var (
		num, soal, skor, sama int
		name, nametemp        string
		soaltemp, skortemp    int = -1, -1
	)

	for {
		status := false
		fmt.Scan(&name)
		if name == "selesai" { //program akan meminta record terus sampai diberi recordan "selesai"
			break
		}
		for i := 0; i < 8; i++ { //tiap orang mengerjakan 8 soal
			fmt.Scan(&num)
			if valid(num) == true {
				fmt.Println("Data yang diEnter tidak valid")
				status = true
			}
			hitungSkor(num, &soal, &skor)
		}
		if status == false {
			if disq(skor) == false {
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
				case (skortemp == skor) && (soaltemp == soal): //jika keduanya sama juga maka keduanya menang
					skortemp = skor
					nametemp = fmt.Sprintf("%s , %s", nametemp, name)
					soaltemp = soal
					sama = 1
				}
			} else {
				fmt.Printf("%s telah didiskualifikasi karena total waktu pengerjaan melebihi 5 hours\n", name)
			}
		}
		soal = 0
		skor = 0
	}
	if skortemp != -1 && soaltemp != -1 {
		if (skortemp != skor || soaltemp != soal) && (sama != 1) {
			fmt.Println(nametemp, soaltemp, skortemp)
		} else {
			fmt.Printf("%s adalah juara bersama %d %d", nametemp, soaltemp, skortemp)
		}
	} else { //jika semua tidak ada yang memenuhi syarat ,maka tidak ada pemenang
		fmt.Println("Tidak ada pemenang.")
	}
}
