package main

import (
	"errors"
	"fmt"
)

func mode1(second int) (int, int, int, error) {
	if second < 0 {
		err := errors.New("Wrong record")
		return 0, 0, 0, err
	}
	hour := 0
	minute := 0
	seconds := 0
	hour = second / 3600
	minute = (second % 3600) / 60
	seconds = (second % 60) % 60
	return hour, minute, seconds, nil
}

func mode2(hour, minute, second int) (int, error) {
	seconds := 0
	if hour < 0 || minute < 0 || second < 0 {
		err := errors.New("Wrong record")
		return 0, err
	}
	seconds = hour*3600 + minute*60 + second
	return seconds, nil
}

// time convertion program
func main() {
	var mode, second, hour, minute int
	fmt.Print("Mode: ")
	fmt.Scan(&mode)
	if mode == 1 {
		fmt.Print("Detik: ")
		fmt.Scan(&second)
		hour, minute, second, err := mode1(second)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("hours: ", hour)
			fmt.Println("Menit: ", minute)
			fmt.Println("Detik: ", second)
		}
	} else if mode == 2 {
		fmt.Print("hours: ")
		fmt.Scan(&hour)
		fmt.Print("Menit: ")
		fmt.Scan(&minute)
		fmt.Print("Detik: ")
		fmt.Scan(&second)
		second, err := mode2(hour, minute, second)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Detik total: ", second)
		}
	} else {
		fmt.Println("Wrong mode")
	}
}
