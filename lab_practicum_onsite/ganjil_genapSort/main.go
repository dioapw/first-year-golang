package main

import "fmt"

//M is a const of lenght of array data
const M = 1000000

//data is a type of array data
type data [M]int

var temp string //to combine result

func fillArray(d *data, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
}

func insertionSort(d *data, n int) {
	for i := 1; i < n; i++ {
		if d[i-1]%2 == 1 { //only odd get sorted first
			j := i
			for j > 0 {
				if d[j-1] > d[j] { //sorted ascending
					d[j-1], d[j] = d[j], d[j-1]
				}
				j--
			}
		}
	}

	for i := 1; i < n; i++ {
		if d[i]%2 == 0 {
			j := i
			for j > 0 {
				if d[j-1] < d[j] { //sorted descending
					d[j-1], d[j] = d[j], d[j-1]
				}
				j--
			}
		}
	}
}

//printArr is a function to combine all the output
func printArr(d data, n int) {
	condition := false
	for j := 0; j < n; j++ {
		if d[j]%2 == 1 { //only odd get output first
			if temp == "" || condition == false { //to make sure the space is right
				temp = fmt.Sprintf("%s%d", temp, d[j])
				condition = true
			} else {
				temp = fmt.Sprintf("%s %d", temp, d[j])
			}
		}
	}
	for i := 0; i < n; i++ {
		if d[i]%2 == 0 {
			if temp == "" || condition == false {
				temp = fmt.Sprintf("%s%d", temp, d[i])
				condition = true
			} else {
				temp = fmt.Sprintf("%s %d", temp, d[i])
			}
		}
	}
}

func main() {
	var num, n int
	var Arr data
	fmt.Scan(&num) //length of input by user
	for i := 0; i < num; i++ {
		fmt.Scan(&n) //length of array by user
		fillArray(&Arr, n)
		insertionSort(&Arr, n)
		printArr(Arr, n)
		temp = fmt.Sprintf("%s%s", temp, "\n") //to create newline each length of array user's
	}
	fmt.Println("-------output--------")
	fmt.Println(temp)
}
