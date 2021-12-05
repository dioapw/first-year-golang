package main

import "fmt"

//func cetakDeret berfungsi untuk memproses deret bilangan
func cetakDeret(n int, sum *int) {
	if n%2 == 0 { //jika deret genap maka dibagi 2
		*sum = n / 2
	} else { //jika ganjil maka 3n+1
		*sum = (3 * n) + 1
	}
}

func main() {
	var num, sum int
	fmt.Scan(&num)
	fmt.Print(num, " ")
	for { //deret bilangan disajikan secara horizontal dan diawali dengan suku awal
		cetakDeret(num, &sum)
		fmt.Print(sum, " ")
		num = sum
		if sum == 1 { //proses percetakan deret akan berhenti jika sudah mencapai angka 1
			break
		}
	}
}
