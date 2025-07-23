package main

import "fmt"

func main() {
	name := "Muhammad"

	switch name {
	case "Habil":
		fmt.Println("Hai Habil")
	case "Budi":
		fmt.Println("Hai Budi")
	default:
		fmt.Println("Hai Boleh Kenalan?")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang ")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// switch tanpa statement mirip if else

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlau panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}