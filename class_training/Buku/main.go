package main

import "fmt"

//penyimpan data-data buku
type buku struct {
	ID        int
	judulBuku string
	penerbit  string
	tahun     int
	genre     string
	harga     int
}

//variabel global
var (
	book = []buku{}
)

//prize func berfungsi menentukan harga buku termahal
func prize(b *buku) {
	sum := book[0].harga
	for i := 1; i < len(book); i++ {
		if sum < book[i].harga {
			sum = book[i].harga
			b.harga = i
		}
	}
}

//program mendata buku
func main() {
	var (
		n   int
		max buku
	)

	fmt.Print("Berapa sum buku yang ingin diEnter: ")
	fmt.Scan(&n)
	book = make([]buku, n) //menentukan batas array

	for i := 0; i < len(book); i++ {
		fmt.Scan(&book[i].ID)
		fmt.Scan(&book[i].judulBuku)
		fmt.Scan(&book[i].penerbit)
		fmt.Scan(&book[i].tahun)
		fmt.Scan(&book[i].genre)
		fmt.Scan(&book[i].harga)
	}

	fmt.Println("Daftar judul buku:")
	for _, x := range book {
		fmt.Println(x.judulBuku)
	}

	prize(&max)
	fmt.Printf("Identitas buku dengan harga termahal: %v", book[max.harga])

}
