package main

import (
	"fmt"
	"strconv"
)

//N adalah constanta dari batas array
const N = 512

//tipe data untuk menyimpan pixel seperti r,g,b
type pixel struct {
	r, g, b int
}

//variabel global
var img [N][N]pixel
var sum [N][N]int

//program untuk membuat ID gambar
func main() {
	var (
		format  string
		w, h    int
		max     int
		r, g, b string
	)
	sumC := 0
	fmt.Scan(&format)
	fmt.Scan(&w, &h) //batas ditentukan berdasarkan user
	fmt.Scan(&max)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Scan(&r, &g, &b)
			//tiap record terdiri dari r,g,b dan berulang terus
			if r != "#" || j != 0 { //jika ada  record "#" di awal baris maka record seluruh baris diabaikan
				sumC += len([]rune(r)) + len([]rune(g)) + len([]rune(b)) //proses menghitung karakter tiap string
				if i1, err := strconv.Atoi(r); err == nil {
					if i1 > max { //ID tidak boleh melebihi max
						sum[i][j] = 0
						break
					}
					sum[i][j] += i1
				}
				if i2, err := strconv.Atoi(g); err == nil {
					if i2 > max { //ID tidak boleh melebihi max
						sum[i][j] = 0
						break
					}
					sum[i][j] += i2
				}
				if i3, err := strconv.Atoi(b); err == nil {
					if i3 > max { //ID tidak boleh melebihi max
						sum[i][j] = 0
						break
					}
					sum[i][j] += i3
				}
				if sumC >= 72 { //jika sum karakter tiap baris melebihi 72 maka tidak dihitung
					for j := 0; j < h; j++ {
						sum[i][j] = 0
					}
					break
				}
			} else {
				sum[i][j] = 0
				break
			}
		}
		sumC = 0
	}
	fmt.Println("P2")
	fmt.Println(w, h)
	fmt.Println(max)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ { //proses enID ID nya
			if sum[i][j] != 0 {
				if sum[i][j]%2 == 0 { //jika sum genap maka akan menjadi 0
					sum[i][j] = 0
					fmt.Print(sum[i][j], " ")
				} else { //jika sum ganjil maka akan menjadi 255
					sum[i][j] = 255
					fmt.Print(sum[i][j], " ")
				}
			}
		}
		fmt.Println()
	}
}
