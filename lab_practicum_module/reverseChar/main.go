package main

import "fmt"

//NMAX is constant array
const NMAX = 127

type tabel [NMAX]string

func main() {
	var m int
	var tab tabel
	fmt.Print("Teks \t\t: ")
	isiArray(&tab, &m)
	bailkanArray(&tab, m)
	fmt.Print("Reverse teks \t: ")
	cetakArray(tab, m)
	fmt.Println()
	fmt.Println("Palindrom ? ", palindrom(tab, m))

}

func isiArray(t *tabel, n *int) {
	*n = 0
	for i := 0; i < len(t); i++ {
		fmt.Scan(&t[i])
		if t[i] == "." {
			break
		}
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
}

func bailkanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}
