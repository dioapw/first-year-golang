package main

import "fmt"

func main() {
	var (
		p1, p2, p3, p4 string
	)
	berhasil := true
	a := 1
	i := 1
	for gagal := false; !gagal; gagal = (p1 != "merah" || p2 != "kuning" || p3 != "hijau" || p4 != "ungu") || i > 5 {
		fmt.Print("Percobaan", a, ":")
		fmt.Scanln(&p1, &p2, &p3, &p4)
		berhasil = p1 == "merah" && p2 == "kuning" && p3 == "hijau" && p4 == "ungu"
		a++
		i++
	}
	for i <= 5 {
		berhasil = false
		fmt.Print("Percobaan", a, ":")
		fmt.Scanln(&p1, &p2, &p3, &p4)
		i++
		a++
	}

	fmt.Println("Berhasil: ", berhasil)
}
