package main

import "fmt"

func main() {
	var (
		cashback, diskon float64
		member           string
		a, b, c, d, e    int
	)

	fmt.Print("Apakah Anda member dan berapa poin Anda: ")
	fmt.Scan(&member, &a, &b, &c, &d, &e)

	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = (float64(a) + float64(b) + float64(c)) * 3.1
		diskon = 0
	} else if a%2 == 1 && b%2 == 1 && c%2 == 1 && d%2 == 1 && e%2 == 1 {
		cashback = 0
		diskon = (float64(c) + float64(d) + float64(e)) * 1.7
	} else {
		cashback = (float64(a) + float64(b) + float64(c)) * 3.1
		diskon = (float64(c) + float64(d) + float64(e)) * 1.7
	}

	if member == "yes" {
		if cashback > 0 {
			cashback = cashback + (cashback * 0.15)
		}
		if diskon > 0 {
			diskon = diskon + (diskon * 0.15)
		}
	}

	if diskon > 50 {
		diskon = 50
	} else if cashback > 35 {
		cashback = 35
	}

	fmt.Printf("Cashback: %.2f\n", cashback)
	fmt.Printf("Diskon: %.2f\n", diskon)
}
