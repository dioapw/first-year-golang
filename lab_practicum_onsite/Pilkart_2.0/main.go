package main

import "fmt"

//MaxMin func berfungsi untuk mencari nilai terbanyak dan terendah
func MaxMin(suara [20]int) {
	max := suara[1]
	max2 := suara[2]
	temp, temp2 := 1, 2
	for i := 2; i < len(suara); i++ {
		if max < suara[i] { //nilai terbanyak menjadi ketua RT
			max = suara[i]
			temp = i
		} else if max2 < suara[i] { //nilai terbanyak kedua menjadi wakil ketua RT
			max2 = suara[i]
			temp2 = i
		}
	}
	fmt.Println("Ketua RT: ", temp)
	fmt.Println("Wakil ketua RT: ", temp2)
}

//program menghitung suara pemilihan RT
func main() {
	suara := [20]int{} //ada 1 - 19 calon
	sah, masuk := 0, 0
	var num int
	for {
		fmt.Scan(&num)
		if num == 0 { //recordan diakhiri angka 0
			break
		}
		if num > 0 && num < 20 { //suara sah jika nomor yang dipilih antara 1-19
			for i := 1; i < len(suara); i++ {
				if num == i {
					suara[i]++
				}
			}
			sah++
		}
		masuk++ //dihitung juga suara yang masuk
	}
	fmt.Println("Suara masuk: ", masuk)
	fmt.Println("Suara sah: ", sah)
	MaxMin(suara) //output siapa pemenangnya
}
