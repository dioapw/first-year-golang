package main

import "fmt"

func main() {
	var (
		nilai, nilai2 int
		name, name2   string
	)
	fmt.Print("Masukan kedua name peserta: ")
	fmt.Scan(&name, &name2)
	berhasil, berhasil1 := 0, 0
	temp := 0
	penentu := 0
	for i := 0; i < 5; i++ {
		penentu = 0
		for j := 1; j <= 3; j++ {
			fmt.Scan(&nilai)
			penentu += nilai
		}
		if penentu == 0 {
			temp++
		} else {
			berhasil++
		}
		penentu = 0
		for j := 1; j <= 3; j++ {
			fmt.Scan(&nilai2)
			penentu += nilai2
		}
		if penentu == 0 {
			temp++
		} else {
			berhasil1++
		}
	}
	if berhasil > berhasil1 {
		fmt.Println(name, "Memenangkan duel dan berhasil menyelesaikan", berhasil, "dari 5 soal")
	} else {
		fmt.Println(name2, "Memenangkan duel dan berhasil menyelesaikan", berhasil1, "dari 5 soal")
	}
}
