package main

import (
	"fmt"
	"math/rand"
	"time"
)

//MAXSIZE is a const of RecType array
const MAXSIZE = 2019 // i use 2019 instead of 20192020 because it's get an error data too big

//RecType is a struct that consist count and size data
type RecType struct {
	count int
	size  int
}

//ArrType is a struct that consist array of RecType
type ArrType [MAXSIZE]RecType

//iSort is a function that sort data count ascending using insertion sort
func iSort(tab *ArrType, nsize int) {
	for i := 1; i < nsize; i++ {
		j := i
		for j > 0 {
			if tab[j-1].count > tab[j].count {
				tab[j-1].count, tab[j].count = tab[j].count, tab[j-1].count
			}
			j--
		}
	}
}

//msort is a function that sort data size descending using selection sort
func mSort(tab *ArrType, nsize int) {
	for i := 0; i < nsize; i++ {
		var maxIdx = i
		for j := i; j < nsize; j++ {
			if tab[j].size > tab[maxIdx].size {
				maxIdx = j
			}
		}
		tab[i].size, tab[maxIdx].size = tab[maxIdx].size, tab[i].size
	}
}

//isFound is a function to search specific data count using binary search
func isFound(tab ArrType, n, v int) bool {
	found := false
	left := 0
	right := n
	for left < right && !found {
		mid := (left + right) / 2
		if tab[mid].count > v {
			right = mid
		} else if tab[mid].count < v {
			left = mid + 1
		} else {
			found = true
		}
	}
	return found
}

//inputCount is a function to fill array data of count
func inputCount(tab *ArrType, nsize int, data []int) {
	for i := 0; i < nsize; i++ {
		tab[i].count = data[i]
		fmt.Print(tab[i].count, " ")
	}
}

//inputCount is a function to fill array data of size
func inputSize(tab *ArrType, nsize int, data []int) {
	for i := 0; i < nsize; i++ {
		tab[i].size = data[i]
		fmt.Print(tab[i].size, " ")
	}
}

func main() {
	var (
		search int
		data   ArrType
	)
	randomCount := generateSlice(20) //to generate 20 random slice
	fmt.Println("--- Count Unsorted --- ")
	inputCount(&data, 20, randomCount) //to fill array of count
	fmt.Println("\n--- Count Sorted using insertion ---")
	iSort(&data, 20) //to sort array of count

	for i := 0; i < 20; i++ {
		fmt.Print(data[i].count, " ") //to output the sorted array
	}

	fmt.Println()
	randomSize := generateSlice(20) //to generate 20 random slice
	fmt.Println("\n--- Size Unsorted --- ")
	inputSize(&data, 20, randomSize) //to fill array of size
	fmt.Println("\n--- Size Sorted using selection --")
	mSort(&data, 20) //to sort array of size

	for i := 0; i < 20; i++ {
		fmt.Print(data[i].size, " ") //to output the sorted array
	}

	fmt.Println()
	fmt.Println()
	fmt.Print("What do you want to find in Count : ")
	fmt.Scan(&search)
	fmt.Println(isFound(data, 20, search)) //to search an integer by user's input in count data
}

//generateSlice function is to generate random slice of interger from -999 to 999
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
