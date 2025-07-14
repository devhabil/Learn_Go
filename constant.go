package main

import "fmt"

func main() {
	/*
	const nilai tidak dapat diubah berbeda dari variable
	*/
	const (
		firstName string = "Muhammad Habil"
		lastName = "Arifin"		
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}