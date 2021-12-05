package main

import "fmt"

//function fact untuk menghitung faktorial
func fact(num int) int {
	sum := 1
	num2 := num
	for i := 0; i < num2; i++ {
		sum *= num
		num--
	}
	return sum
}

//func perm untuk menghitung permutasi
func perm(n int) int {
	var n2, n3 int
	sum := 1
	sum = fact(n)
	fmt.Scan(&n2)
	if n2 > n {
		n2 = n
	}
	sum = sum / fact(n2)
	for n2 < n { //program berhenti apabila selisih bilangan kedua dan seterusnya lebih besar dari bilangan pertama
		fmt.Scan(&n3)
		if n2+n3 > n { //jika ada bilangan yang membuat sumnya lebih besar dari bilangan pertama maka hanya diambil selisihnya
			n3 = n - n2
		}
		sum = sum / fact(n3)
		n2 += n3
	}
	return sum
}

//program menghitung permutasi
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(perm(n))
}
