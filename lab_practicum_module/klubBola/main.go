package main

import "fmt"

var hasil []string

func main() {
	var (
		nameA, nameB   string
		gamesA, gamesB int
	)
	fmt.Print("Klub A : ")
	fmt.Scan(&nameA)
	fmt.Print("Klub B : ")
	fmt.Scan(&nameB)
	i := 0
	for {
		fmt.Print("Pertandingan ", i+1, " : ")
		fmt.Scan(&gamesA, &gamesB)
		if gamesA == -1 || gamesB == -1 {
			break
		}
		if gamesA > gamesB {
			hasil = append(hasil, nameA)
		} else if gamesA < gamesB {
			hasil = append(hasil, nameB)
		} else {
			hasil = append(hasil, "Draw")
		}
		i++
	}

	for i, val := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, val)
	}
	fmt.Println("Pertandingan selesai")

}
