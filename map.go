package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Eko"
	//person["address"] = "Subang"

	// atau langsung 
	person := map[string]string{
		"name": "Habil",
		"address": "Kemiling",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Habil"
	book["ups"] = "Salah"

	fmt.Println(book)
	
	delete(book, "ups")

	fmt.Println(book)
}