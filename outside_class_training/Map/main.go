package main

import "fmt"

func main() {
	//define map
	email := make(map[string]string)

	//assign kv
	email["Bobs"] = "bobs@gmail.com"
	email["Sharon"] = "sharon@gmail.com"
	email["Mike"] = "mike@gmail.com"

	fmt.Println(email)
	fmt.Println(len(email))
	fmt.Println(email["Bobs"])

	//looping

	for key, val := range email {
		fmt.Println(key, "\t:", val)
	}

	for key := range email {
		fmt.Printf("Name: %s \n", key)
	}

	//delete
	delete(email, "Bobs")
	fmt.Println(email)

	//define & add kv

	phones := map[string]string{"dio": "01928291", "yudha": "081208017", "rezki": "0981260972"}

	fmt.Println(phones)

	//existance checking

	chicken := map[string]int{"januari": 50, "februari": 40}
	value, isExist := chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	//combien

	chickens := []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
