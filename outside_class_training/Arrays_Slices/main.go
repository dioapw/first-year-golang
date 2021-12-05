package main

import "fmt"

func main() {

	var fruit [3]string
	fruit[0] = "Apple"
	fruit[1] = "Grape"
	fruit[2] = "Orange"

	fmt.Println(fruit)
	fmt.Println(fruit[1])

	// different ways

	fruits := []string{"Apple", "Orange", "Grape", "mango"}

	fmt.Println(fruits)
	fmt.Println(fruits[2])

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	//append & slices

	var cFruits = append(aFruits, "papaya")
	fmt.Println(cFruits)

	fruits = append(fruits, "Banana")
	fmt.Println(fruits)
	// slices

	fruitslices := []string{"Apple", "Orange", "Grape", "cherry"}

	fmt.Println("jumlag data: ", len(fruitslices))
	fmt.Println("macam data: \t", fruitslices[1:3])

	//other

	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers2", numbers2)

	//for - range

	var fruits2 = [...]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits2 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	for _, fruit := range fruits2 {
		fmt.Printf("name buah : %s\n", fruit)
	}

	//copy

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pinnaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2) // potato potato potato
	fmt.Println(src2) // watermelon pinnaple
	fmt.Println(n2)   // 2

	// 3 indeks

	var fruit2 = []string{"apple", "grape", "banana"}
	var aFruits2 = fruit2[0:2]
	var bFruits2 = fruit2[0:2:2]

	fmt.Println(fruit2)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruit2)) // len: 3
	fmt.Println(cap(fruit2)) // cap: 3

	fmt.Println(aFruits2)      // ["apple", "grape"]
	fmt.Println(len(aFruits2)) // len: 2
	fmt.Println(cap(aFruits2)) // cap: 3

	fmt.Println(bFruits2)      // ["apple", "grape"]
	fmt.Println(len(bFruits2)) // len: 2
	fmt.Println(cap(bFruits2)) // cap: 2
}
