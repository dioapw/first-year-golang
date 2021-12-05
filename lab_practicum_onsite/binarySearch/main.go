package main

import "fmt"

//NMAX is const of array data
const NMAX = 1000000

var data [NMAX]int //data is an array of integer

func main() {
	var N, K int
	fmt.Scan(&N, &K) //to input len & value to find
	isiArray(N)
	if binarySearch(K, N) != -1 { //to check if its found it will output its index
		fmt.Println(binarySearch(K, N))
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
func binarySearch(k int, n int) int {
	low := 0
	high := n
	for low <= high {
		median := (low + high) / 2
		if data[median] < k {
			low = median + 1
		} else if data[median] > k {
			high = median - 1
		} else {
			return median
		}
	}
	return -1
}
