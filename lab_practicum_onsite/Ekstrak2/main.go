package main

import "fmt"

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
		format string
		w, h   int
		max    int
		ID   int
	)
	count := 1
	temp := 0
	fmt.Scan(&format)
	fmt.Scan(&w, &h) //batas ditentukan berdasarkan user
	fmt.Scan(&max)

	for i := 0; i < w; i++ {
		for j := 0; j < h*3; j++ {
			fmt.Scan(&ID)

			if count > 3 {
				count = 1
				temp = 0
			}
			if ID < max { //ID tidak boleh melebihi max
				switch { //tiap record terdiri dari r,g,b dan berulang terus
				case count == 1:
					img[i][j].r = ID
					temp += img[i][j].r
				case count == 2:
					img[i][j].g = ID
					temp += img[i][j].g
				default:
					img[i][j].b = ID
					temp += img[i][j].b
				}
			} else {
				sum[i][j] = 0
				break
			}
			if count == 3 { //menghitung hasilnya tiap rgb
				sum[i][j] = temp
			}
			count++
		}
	}
	fmt.Println("P2")
	fmt.Println(w, h)
	fmt.Println(max)
	for i := 0; i < w; i++ {
		for j := 0; j < h*3; j++ { //proses enID ID nya
			if count > 3 {
				count = 1
			}
			if count == 3 && sum[i][j] != 0 {
				if sum[i][j]%2 == 0 { //jika sum genap maka akan menjadi 0
					sum[i][j] = 0
					fmt.Print(sum[i][j], " ")
				} else { //jika sum ganjil maka akan menjadi 255
					sum[i][j] = 255
					fmt.Print(sum[i][j], " ")
				}
			}
			count++
		}
		fmt.Println()
	}
}
