package main

import (
	"fmt"
	"strconv"
)

type person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting method (value receiver)
func (p person) Greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *person) hasBirthday() {
	p.age++
}

//getMarried method (pointer receiver)
func (p *person) getMarried(spouselastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouselastName
	}
}
func main() {
	// init person using struct
	person1 := person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F", age: 25}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	//alternate
	person2 := person{"Bob", "Johnson", "New York", "M", 30}

	// function
	fmt.Println(person1.Greet())
	person1.hasBirthday()
	person1.getMarried("Willliam")
	fmt.Println(person1.Greet())

	fmt.Println(person2.Greet())
	person2.hasBirthday()
	person2.getMarried("De Santa")
	fmt.Println(person2.Greet())
}
