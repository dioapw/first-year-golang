package main

import "fmt"

func main() {
	var record string
	sum, sumf := 0, 0
	fmt.Print("record: ")
	fmt.Scan(&record) // program will always loop if its given record not "q"
	if record != "q" {
		for stop := false; !stop; {
			if record == "w" {
				sum = 3
			} else if record == "d" {
				sum = 1
			} else if record == "l" {
				sum = 0
			}
			sumf += sum
			sum = 0
			fmt.Print("record: ")
			fmt.Scan(&record)
			stop = (record == "q")
		}
	}
	// program will output the sum of all the record
	fmt.Println("sum poin yang diperoleh adalah ", sumf)
}
