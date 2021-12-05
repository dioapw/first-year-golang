package main

import (
	"fmt"
)

func main() {
	var a1, a2, a3, b1, b2, b3 int
	var absolutAhmad, absolutBadrun bool
	var menangAhmad, menangBadrun bool

	fmt.Scanln(&a1, &a2, &a3)
	fmt.Scanln(&b1, &b2, &b3)

	absolutAhmad = ((a1 > b1) && (a2 > b2) && (a3 > b3))
	absolutBadrun = ((b1 > a1) && (b2 > a2) && (b3 > a3))

	menangAhmad =
		((a1 > b1) && (a2 > b2)) ||
			((a2 > b2) && (a3 > b3)) ||
			((a1 > b1) && (a3 > b3)) ||
			((a1 > b1) && (a2 == b2) && (a3 == b3)) ||
			((a2 > b2) && (a1 == b1) && (a3 == b3)) ||
			((a3 > b3) && (a1 == b1) && (a2 == b2))

	menangBadrun =
		((b1 > a1) && (b2 > a2)) ||
			((b2 > a2) && (b3 > a3)) ||
			((b1 > a1) && (b3 > a3)) ||
			((b1 > a1) && (a2 == b2) && (a3 == b3)) ||
			((b2 > a2) && (a1 == b1) && (a3 == b3)) ||
			((b3 > a3) && (a1 == b1) && (a2 == b2))

	fmt.Printf("Pemenang absolut -> Ahmad? %t, Badrun? %t\n", absolutAhmad, absolutBadrun)
	fmt.Printf("Pemenang kompetisi -> Ahmad? %t, Badrun? %t\n", menangAhmad, menangBadrun)
}
