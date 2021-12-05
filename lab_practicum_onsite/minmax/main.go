package main

import "fmt"

//N is a const for array
const N = 2000

//array is a type of array integer
type array [N]int

func main() {
	var N, m, frek int
	var data array
	sum := ""
	fmt.Scan(&N) //input how many data to input
	for i := 0; i < N; i++ {
		isiArray(&data, &m)                                                        //to fill array
		hitungFrekuensi(data, m, min(data, m), max(data, m), &frek)                //to count frequency of max number
		sum = fmt.Sprintf("%s %d %d %d \n", sum, min(data, m), max(data, m), frek) //to keep the value until the process finished
	}
	fmt.Print(sum)
}

//isiArray is a function to fill an array
func isiArray(data *array, m *int) {
	*m = 0
	for i := 0; i < len(data); i++ {
		fmt.Scan(&data[*m])
		if data[*m] == 999 {
			data[*m] = 0
			break
		}
		*m++
	}
}

//max is a searching function to find maximum number
func max(data array, m int) int {
	maxValue := data[0]
	for i := 1; i < m; i++ {
		if maxValue < data[i] {
			maxValue = data[i]
		}
	}
	return maxValue
}

func min(data array, m int) int {
	minValue := data[0]
	for i := 1; i < m; i++ {
		if minValue > data[i] {
			minValue = data[i]
		}
	}
	return minValue
}

//hitungFrekuensi is function to count how many frequency of max number
func hitungFrekuensi(data array, m int, key1 int, key2 int, frek *int) {
	*frek = 0
	for i := 0; i < m; i++ {
		if key1 != data[i] && data[i] != key2 {
			*frek++
		}
	}
}
