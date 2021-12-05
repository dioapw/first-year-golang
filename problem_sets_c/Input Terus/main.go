package main

import "fmt"

func main() {
	var C byte

	Text := ""
	fmt.Scanf("%c\n", &C)

	for C != '-' {
		Text = fmt.Sprintf("%s%c", Text, C)
		fmt.Scanf("%c\n", &C)
	}

	fmt.Println(Text)
}
