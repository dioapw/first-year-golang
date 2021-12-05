package main

import "fmt"

//M is a const of lenght of array data
const M = 1000000

//data is a type of array data
var data [M]int

func insertionSort(n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if data[j-1] < data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j--
		}
	}

	if n%2 == 1 {
		fmt.Print(data[n/2], " ")
	} else {
		median := n / 2
		fmt.Print((data[median]+data[median-1])/2, " ")
	}
}

func main() {
	var num int
	i := 0
	for stop := false; !stop; {
		fmt.Scan(&num)
		if num == 0 {
			insertionSort(i)
			i = 0
		} else {
			data[i] = num
			i++
		}
		stop = num == -59672
	}

}
