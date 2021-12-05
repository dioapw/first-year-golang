package main

import "fmt"

func main() {
	var nilai int
	sum := 0
	sumI := 0
	kelipatan := 0
	i := 1

	for stop := false; !stop; {
		fmt.Scanln(&nilai)
		if nilai%i == 0 {
			kelipatan++
			sum += nilai
			sumI += i
		}
		stop = (kelipatan == 5)
		i++
	}

	fmt.Println("sum bilangan: ", sum)
	fmt.Println("sum posisi  : ", sumI)
}
