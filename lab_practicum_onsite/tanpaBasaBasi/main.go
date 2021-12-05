package main

import "fmt"

//NMAX is const of array data
const NMAX = 1000000

var data [NMAX]int //data is an array of integer

func main() {
	var N, K int
	fmt.Scan(&N, &K) //to input len & value to find
	isiArray(N)
	if posisi(N, K) != -1 { //to check if its found it will output its index
		fmt.Println(posisi(N, K))
	} else { //if not
		fmt.Println("TIDAK ADA")
	}
}

//isiArray is a function to fill data array
func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

//posisi is a searching function
func posisi(n, k int) int {
	for i := 0; i < n; i++ {
		if k == data[i] {
			return i //found
		}
	}
	return -1 //not found
}
