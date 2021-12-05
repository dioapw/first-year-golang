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
	var num int
	temp := 0
	sum := 0
outerloop:
	for {
		fmt.Scanln(&num)
		if num < 1000000 { //recordan harus diatas 1000000
			for {
				fmt.Scanln(&num)
				if num > 1000000 {
					break
				}
			}
		}
		temp = num
		for {
			cetakDeret(num, &sum)
			num = sum
			if sum == 1 { //proses deret berhenti sampai angkanya adalah 1
				fmt.Printf("Deret yang diawali %d pasti akan berhenti\n", temp)
				break outerloop
			} else if sum < 1 { //proses tidak berhenti
				fmt.Printf("Deret yang diawali %d tidak akan berhenti\n", temp)
				break outerloop
			}
		}
	}
}
