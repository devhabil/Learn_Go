package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Habil"
	names[2] = "Arifin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		80,
		90,
		95,
		110,

	}

	fmt.Println(values)		
	fmt.Println(len(values))
}