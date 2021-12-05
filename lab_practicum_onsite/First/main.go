package main

import "fmt"

func main() {

	var name string
	name = "DIO ADITYA"

	age := 18

	var address string
	var angka int
	var desimal float64
	var kondisi bool

	angka = 69
	desimal = 6.344242
	kondisi = false

	fmt.Println(angka)
	fmt.Println(desimal)
	fmt.Println(kondisi)

	fmt.Println(name, " ", age)
	fmt.Println("DIO ADITYA")
	fmt.Println("J'AI 18 AN")
	fmt.Println("JE SUIS INDONESIEN")
	fmt.Println("JE SUIS FRANCHOPHONE")

	fmt.Println("PRENOM : ")
	fmt.Scan(&name)
	fmt.Println("BONJOUR", name)

	fmt.Println("TOLONG ISI BIODATA!")
	fmt.Println("name saya = ")
	fmt.Scan(&name)
	fmt.Println("age saya = ")
	fmt.Scan(&age)
	fmt.Println("Saya tinggal di = ")
	fmt.Scan(&address)

	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Println("address = ", address)
}
