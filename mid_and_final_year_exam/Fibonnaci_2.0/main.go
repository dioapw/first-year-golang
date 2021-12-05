package main

import "fmt"

func main() {
	var record int
	a := 0
	b := 1
	c := 1
	for gagal := false; !gagal; { //jika record -1 maka akan berhenti
		fmt.Scanln(&record)
		for i := 0; i < record; i++ {
			a = b
			b = c
			c = b + a
		}
		for i := 1; record != -1 && i < 2 && record != 0; i++ { //hasil tiap output adalah fibonnaci terakhir dari record
			fmt.Println(b)
		}
		for j := 1; record == 0 && j < 2; j++ { //jika record 0 maka output 0
			fmt.Println(a)
		}
		gagal = record == -1
		a = 0
		b = 1
		c = 1
	}
	fmt.Println("SELESAI")
}
