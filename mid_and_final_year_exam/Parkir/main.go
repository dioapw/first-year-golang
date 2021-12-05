package main

import (
	"errors"
	"fmt"
)

// parking tariff program
func time(hour1, hour2 int) (int, error) {
	sum := 0
	tariff := 0

	if hour1 == hour2 {
		err1 := errors.New("ERROR , WRONG record")
		return 0, err1
	}

	if hour2 > hour1 {
		sum = hour2 - hour1
	} else {
		sum = (12 - hour1) + hour2
	}

	if sum-2 > 0 { // every first 2 hours it will cost Rp 5000 and next hours after it will cost Rp 3000
		tariff = 5000 + ((sum - 2) * 3000)
	} else if sum-2 == 0 {
		tariff = 5000
	} else { //if the difference between those two hours is one or zero its cost will be free
		err1 := errors.New("Parkir gratis")
		return 0, err1
	}
	return tariff, nil
}

func main() {
	var hour1 int
	var hour2 int

	fmt.Print("hours Parkir: ")
	fmt.Scan(&hour1)
	fmt.Print("hours keluar: ")
	fmt.Scan(&hour2)

	cost, err1 := time(hour1, hour2)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Tarif Rp. %d \n", cost)
	}
}
