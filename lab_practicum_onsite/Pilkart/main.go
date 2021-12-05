package main

import "fmt"

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
	for i := 1; i < len(suara); i++ { //mengouputkan sum suara tiap calon yang mendapatkan suara
		if suara[i] != 0 {
			fmt.Printf("%d : %d\n", i, suara[i])
		}
	}
}
