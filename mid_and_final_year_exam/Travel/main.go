package main

import "fmt"

func sumCost(place, duration, budget int) (int, int, int, int) {
	sumTiket, sumLama, sumMakan, sumDollar := 0, 0, 0, 0
	sumTiket = 13 * place
	sumLama = 514 * duration
	sumMakan = (duration * 2) * 20
	sumDollar = sumLama + sumMakan + sumTiket + 450
	return sumTiket, sumLama, sumMakan, sumDollar

}

func convertion(sumDollar, budget int) (int, int, int) {
	moreLess, moreLess2 := 0, 0
	sumRupiah, sumRupiah2 := 0, 0
	sumRupiah = sumDollar * 13000
	sumRupiah2 = budget * 13000
	moreLess = budget - sumDollar
	moreLess2 = sumRupiah2 - sumRupiah
	return sumRupiah, moreLess, moreLess2
}

func main() {
	var (
		lama   int
		tempat int
		budget int
	)

	fmt.Print("Enter budget dalam dolar : ")
	fmt.Scan(&budget)
	fmt.Print("Masukan sum tempat wisata yang dikunjungi : ")
	fmt.Scan(&tempat)
	fmt.Print("Lama menginap: ")
	fmt.Scan(&lama)

	tiket, lama, makan, dollar := sumCost(tempat, lama, budget)

	fmt.Println()
	fmt.Println("Biaya tiket masuk: $", tiket)
	fmt.Println("Biaya penginapan: $", lama)
	fmt.Println("Biaya konsumsi: $", makan)
	fmt.Println("Total biaya perjalanan(dalam dolar): $", dollar)

	rupiah, kdollar, krupiah := convertion(dollar, budget)

	fmt.Println("Total biaya perjalanan(dalam rupiah): Rp", rupiah)
	fmt.Println("Total kekurangan/kelebihan (dalam dolar) : $", kdollar)
	fmt.Println("Total kekurangan/kelebihan (dalam rupiah): Rp", krupiah)
}
