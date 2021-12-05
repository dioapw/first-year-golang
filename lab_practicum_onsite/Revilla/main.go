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

//func bongkarDeret untuk menghitung sum suku di deret tersebut dan nilai tertinggi nya
func bongkarDeret(n int, tot, max *int) {
	if *max < n {
		*max = n
	}
	*tot = *tot + 1
}

func main() {
	var (
		num, sum int
		temp     int
		tot, max int
	)
	z := ""
	for {
		fmt.Scanln(&num)
		if num < 0 || num > 1000000 { //recordan harus bilangan positif dan dibawah 1000000
			for {
				fmt.Scanln(&num)
				if num > 0 && num < 1000000 {
					break
				}
			}
		}
		temp = num
		if num == 0 { //program akan terus meminta recordan sampai diberi recordan 0
			break
		}
		bongkarDeret(num, &tot, &max)
		for {
			cetakDeret(num, &sum)
			bongkarDeret(sum, &tot, &max)
			num = sum
			if sum == 1 { //proses deret berhenti sampai angkanya adalah 1
				break
			}
		}
		z = fmt.Sprintf("%s%d %d %d\n", z, temp, max, tot) //untuk menyimpan hasil output
		tot = 0
		max = 0
	}
	fmt.Println()
	fmt.Print(z) //output terdiri dari suku awal , nilai suku tertinggi, dan sum suku
}
