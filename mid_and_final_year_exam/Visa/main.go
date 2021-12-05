package main

import "fmt"

//program menghitung kapan pembuatan visa jadi
func tahunKabisat(year int) bool { //jika tahun tersebut kabisat
	status := false
	status = year%4 == 0 && year%400 == 0 || year%100 != 0 && year%4 == 0
	return status
}

func sumDay(month, year int) int { // hari tiap month
	day := 0
	switch month {
	case 1:
		day = 31
	case 2:
		if tahunKabisat(year) == true {
			day = 29
		} else {
			day = 28
		}
	case 3:
		day = 31
	case 4:
		day = 30
	case 5:
		day = 31
	case 6:
		day = 30
	case 7:
		day = 31
	case 8:
		day = 31
	case 9:
		day = 30
	case 10:
		day = 31
	case 11:
		day = 30
	case 12:
		day = 31
	}
	return day
}

func getDate(date, year, month int, day string) (int, int, int, string) { // penentu tanggal,month,tahun visa nya
	time := 0
	monthDay := 0
	month2 := 0
	date2 := 0
	year2 := 0
	if day == "kamis" || day == "jumat" { // karena hari kerja bukan sabtu atau minggu jadi dilewatkan
		time = 4
	} else { //pengurusan visa dilakukan 2 hari
		time = 2
	}
	date2 = date + time            //tanggal jadi
	monthDay = sumDay(month, year) //hari dimonth tersebut
	if date2 <= monthDay {         // jika hari tidak mendekati akhir month
		month2 = month
		year2 = year
	} else { //hari mendekati akhir month maka bisa bertambah tahun atau bertambah month
		date2 = date2 - monthDay //tanggal jadi
		if month == 12 {         //jika bertambah tahun month kembali ke 1
			year2 = year + 1
			month2 = 1
		} else { //jika tidak maka hanya month yang bertambah
			month2 = month + 1
			year2 = year
		}
	}
	return date2, year2, month2, day
}

func main() {
	var (
		date  int
		year  int
		month int
		day   string
	)
	fmt.Print("Masukan tanggal, tahun, month dan hari secara berurutan : ")
	fmt.Scan(&date, &month, &year, &day)
	whatDate, whatYear, whatMonth, whatDay := getDate(date, year, month, day) // harus berurutan dengan record function
	whatDay = ""                                                              // karena want dan return harus sama di function sedangkan hanya membutuhkan 3 output dar 4 record sehingga hari ditiadakan
	fmt.Printf("Tanggal: %d , month: %d , Tahun: %d%s\n", whatDate, whatMonth, whatYear, whatDay)

}
