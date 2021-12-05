package main

import "fmt"

func main() {
	var (
		x, nilai, nilai2 int
		name, name2      string
	)
	fmt.Print("Masukan sum soal dan kedua name peserta: ")
	fmt.Scan(&x, &name, &name2)
	berhasil, berhasil1 := 0, 0
	temp := 0
	penentu := 0
	fmt.Println("silahkan", name, "meEnter nilai: ")
	for i := 0; i < x; i++ {
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
	}
	fmt.Println("silahkan", name2, "meEnter nilai: ")
	for i := 0; i < x; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Scan(&nilai2)
			penentu += nilai2
		}
		if penentu == 0 {
			temp++
		} else {
			berhasil1++
		}
		penentu = 0
	}
	if berhasil > berhasil1 {
		fmt.Println(name, "Memenangkan duel dan berhasil menyelesaikan", berhasil, "dari", x, "soal")
	} else {
		fmt.Println(name2, "Memenangkan duel dan berhasil menyelesaikan", berhasil1, "dari", x, "soal")
	}
}
