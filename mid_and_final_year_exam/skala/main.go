package main

import "fmt"

func main() {
	var n, permenawal, selisih, permen int

	fmt.Print("Berapa child: ")
	fmt.Scan(&n)

	jumpermen := 0
	permenakhir := 0

	for i := 1; i <= n; i++ {
		fmt.Print("child ke ", i, " : ")
		fmt.Scan(&permenawal)

		selisih = permenawal - permenakhir 

		if permenawal > permenakhir {
			permen = 1 + selisih // jika permen awal lebih besar dari permen akhir maka ditambah selisih
		} else {
			permen = 1
		}

		if i == 1 {
			permen = 1 // minimal permen adalah 1
		}

		permenakhir = permenawal
		jumpermen += permen

	}
	fmt.Println("Minimum permen adalah", n)
	fmt.Println("sum permen yg diberikan", jumpermen)
}
