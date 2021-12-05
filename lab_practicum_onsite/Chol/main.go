package main

import "fmt"

func main() {
	var hdl, ldl, trly int
	total := float64(0)
	fmt.Print("HDL: ")
	fmt.Scan(&hdl)
	fmt.Print("LDL: ")
	fmt.Scan(&ldl)
	fmt.Print("Triglyserida: ")
	fmt.Scan(&trly)

	total = (0.2 * float64(trly))
	total = total + float64(hdl) + float64(ldl)

	fmt.Println("Total: ", total)
	fmt.Println("Rasio total/HDL: ", total/float64(hdl))
	fmt.Println("Sehat? ", total/float64(hdl) < 5)
}
