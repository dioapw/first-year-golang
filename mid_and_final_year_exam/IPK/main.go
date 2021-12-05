package main

import "fmt"

func main() {
	var ipk int
	var sumA float64 = 0
	var sumB float64 = 0

	countA := 0
	countB := 0
	i := 1

	fmt.Print("Mahasiswa ", i, " : ")
	fmt.Scan(&ipk)
	for ipk != -1 {
		if countA+5 < countB {
			sumA += float64(ipk)
			countA++
			fmt.Println("Mahasiswa ", i, " nomor ", countA, " ruang A")
		} else if countB+5 < countA {
			sumB += float64(ipk)
			countB++
			fmt.Println("Mahasiswa ", i, " nomor ", countB, " ruang B")
		} else if sumA <= sumB {
			sumA += float64(ipk)
			countA++
			fmt.Println("Mahasiswa ", i, " nomor ", countA, " ruang A")
		} else {
			sumB += float64(ipk)
			countB++
			fmt.Println("Mahasiswa ", i, " nomor ", countB, " ruang B")
		}
		i++
		fmt.Print("Mahasiswa ", i, " : ")
		fmt.Scan(&ipk)
	}
	fmt.Println("Ruang A sum: ", countA, "Rerata IPK: ", sumA/float64(countA))
	fmt.Println("Ruang B sum: ", countB, "Rerata IPK: ", sumB/float64(countB))
}
