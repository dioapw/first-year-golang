package main

import "fmt"

func main() {
	var bil1, bil2, bil3, fak int

	fmt.Print("Masukan 3 bilangan: ")
	fmt.Scanln(&bil1, &bil2, &bil3)
	for bil1 > 0 && bil2 > 0 && bil3 > 0 {
		if bil1 < bil2 && bil1 < bil3 {
			fak = bil1
		} else if bil2 < bil1 && bil2 < bil3 {
			fak = bil2
		} else {
			fak = bil3
		}

		found := false
		for !found {
			if bil1%fak == 0 && bil2%fak == 0 && bil3%fak == 0 {
				found = true
			} else {
				fak = fak - 1
			}
		}

		fmt.Println("FPB dari ketiga bilangan: ", fak)
		fmt.Print("Masukan 3 bilangan: ")
		fmt.Scanln(&bil1, &bil2, &bil3)
	}
	if bil1 == 0 && bil2 == 0 && bil3 == 0 {
		fmt.Println("Selesai")
	}
}
