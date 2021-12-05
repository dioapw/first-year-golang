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

//program untuk membuat ID gambar
func main() {
	var (
		format string
		w, h   int
		max    int
		ID   int
	)

	count := 1
	fmt.Scan(&format)
	fmt.Scan(&w, &h) //batas ditentukan berdasarkan user
	fmt.Scan(&max)

	for i := 0; i < w; i++ {
		for j := 0; j < h*3; j++ {
			fmt.Scan(&ID)

			if count > 3 {
				count = 1
			}
			if ID < max { //ID tidak boleh melebihi max
				switch { //tiap record terdiri dari r,g,b dan berulang terus
				case count == 1:
					img[i][j].r = ID
				case count == 2:
					img[i][j].g = ID
				default:
					img[i][j].b = ID

				}
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
			if count == 1 && img[i][j].r != 0 {
				if img[i][j].r%2 == 0 { //jika ID genap maka akan menjadi 0
					img[i][j].r = 0
					fmt.Print(img[i][j].r, " ")
				} else { //jika ID ganjil maka akan menjadi 255
					img[i][j].r = 255
					fmt.Print(img[i][j].r, " ")
				}
			}
			count++
		}
		fmt.Println()
	}
}
