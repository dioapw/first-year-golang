package main

import "fmt"

func main() {
	var weight int
	var bMin, bMax float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&weight)
	arrBerat := make([]float64, weight)
	for i := 0; i < len(arrBerat); i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&arrBerat[i])
	}
	hitungMinMax(arrBerat, &bMin, &bMax)
	fmt.Printf("Berat balita minimum: %.2f Kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f Kg\n", bMax)
	fmt.Printf("Rerata berat balita : %.2f Kg\n", rerata(arrBerat))

}

func hitungMinMax(arrBerat []float64, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for _, val := range arrBerat {
		if *bMax < val {
			*bMax = val
		} else if *bMin > val {
			*bMin = val
		}
	}
}

func rerata(arrBerat []float64) float64 {
	var sum float64
	for _, val := range arrBerat {
		sum += val
	}
	return sum / float64(len(arrBerat))
}
