package main

import "fmt"

func main() {
	var record string
	fmt.Print("record: ")
	fmt.Scan(&record) // a program that will always loop every record its given except "w" / "d" / "l"
	if record != "w" && record != "d" && record != "l" {
		for stop := false; !stop; {
			fmt.Print("record: ")
			fmt.Scan(&record)
			stop = (record == "w" || record == "d" || record == "l")
		}
	}
	if record == "w" { // every "w" record its will output 3 point
		record = "3"
	} else if record == "d" { // every "d" record its will output 1 point
		record = "1"
	} else { // every "l" record its will output 0 point
		record = "0"
	}
	fmt.Println("Poin: ", record)
}
