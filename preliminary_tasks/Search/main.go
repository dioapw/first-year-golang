package main

import "fmt"

//N is const of Rectype array
const N = 2019

//RecType is a struct data that consist 3 var
type RecType struct {
	f1 string
	f2 int
	f3 float64
}

//ArrType is a struct data that consist an array of RecType
type ArrType [N]RecType

//rmax function is to find the max value of real data type
func rmax(data ArrType) float64 {
	max := data[0].f3
	for i := 0; i < len(data); i++ {
		if data[i].f3 > max {
			max = data[i].f3
		}
	}
	return max
}

//imin function is to find the index minimum value of integer type
func imin(data ArrType) int {
	min := data[0].f2
	temp := 0
	for i := 0; i < len(data); i++ {
		if data[i].f2 < min {
			min = data[i].f2
			temp = i
		}
	}
	return temp
}

//found function is to search specific string data
func found(data ArrType, key string) bool {
	for i := 0; i < len(data); i++ {
		if key == data[i].f1 {
			return true
		}
	}
	return false
}

//pos function is to search index specific string data using binary search
func pos(data ArrType, key string) int {
	var searchCount int
	found := false
	left := 0
	right := N
	for left < right && !found {
		mid := (left + right) / 2
		if data[mid].f1 > key {
			right = mid
		} else if data[mid].f1 < key {
			left = mid + 1
		} else {
			searchCount = mid
			found = true
		}
	}
	if found == false { //if the data not found it will output -1
		searchCount = -1
	}
	return searchCount
}

func main() {
	var key string
	status := true
	count := 0
	Arr := ArrType{}
	fmt.Print("record string , integer , dan real: ")
	for i := 0; i < len(Arr); i++ {
		fmt.Scan(&Arr[i].f1, &Arr[i].f2, &Arr[i].f3)
	}
	//to check if the array is sorted Ascending
	for i := 0; i < len(Arr)-1; i++ {
		if Arr[i].f1 < Arr[i+1].f1 {
			count++
		}
	}
	if count == len(Arr)-1 { //if the array is sorted status will have the value of false
		status = false
	}
	//if the array is full
	if len(Arr) == N {
		fmt.Println("Nilai terbesar bilangan real: ", rmax(Arr))
		fmt.Println("Index dengan nilai integer terkecil: ", imin(Arr))
		fmt.Print("string yang ingin dicari: ")
		fmt.Scan(&key)
		if found(Arr, key) == true {
			fmt.Println("Data is found")
		} else {
			fmt.Println("Data is not found")
		}
		if status == false { //if array is sorted
			fmt.Println("Data sudah terurut")
			fmt.Print("string yang ingin dicari: ")
			fmt.Scan(&key)
			fmt.Println("index string yang ingin dicari: ", pos(Arr, key))
		} else {
			fmt.Println("Data belum terurut tidak dapat melakukan binary search")
		}
	}
}
