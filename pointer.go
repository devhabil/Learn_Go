package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Lampung", "Metro", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Bandar Lampung"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi Bandar Lampung

	var address1 Address = Address{"Lampung", "Metro", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandar Lampung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandar Lampung
}