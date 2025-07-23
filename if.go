package main

import "fmt"

func main() {
	name := "Muhammad"
	
	if name == "Habil" {
		fmt.Println("Hello Habil")
	} else if name =="Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// kode if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Besar")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}