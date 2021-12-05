package main

import "fmt"

func getSum(R, G, B float64) (float64, float64, float64) {
	y := 0.0
	cb := 0.0
	cr := 0.0
	y = 16 + ((65.738 * float64(R)) / 256) + ((129.057 * float64(G)) / 256) + ((25.064 * float64(B)) / 256)
	cb = 128 - ((37.945 * float64(R)) / 256) - ((74.494 * float64(G)) / 256) + ((112.439 * float64(B)) / 256)
	cr = 128 + ((112.439 * float64(R)) / 256) - ((94.154 * float64(G)) / 256) - ((18.285 * float64(B)) / 256)
	return y, cb, cr
}

func main() {
	var R, G, B float64

	fmt.Print("Nilai RGB: ")
	fmt.Scan(&R, &G, &B)

	sumY, sumCB, sumCR := getSum(R, G, B)
	fmt.Printf("Y: %.3f, Cb: %.3f, dan Cr: %.3f\n", sumY, sumCB, sumCR)
}
