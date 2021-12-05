package main

import "fmt"

func main() {
	var eprt, n int
	rata := float64(0)
	sum := float64(0)
	j := 1
	fmt.Print("Jumalh calon wisudawan: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("calon ", i, " eprt ", j, ": ")
		fmt.Scanln(&eprt)
		for eprt < 500 && j < 3 {
			j++
			fmt.Print("calon ", i, " eprt ", j, ": ")
			fmt.Scanln(&eprt)
		}
		sum += float64(eprt)
		j = 1
	}
	rata = sum / float64(n)
	fmt.Printf("Rerata eprt: %.1f", rata)
}
