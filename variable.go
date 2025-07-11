package main

import "fmt"

func main() {
	/*
		Digolang tidak wajib menyebutkan tipe data variablenya
		:= simbol disamping agar kita tidak menggunakan kata var, ini boleh dipakai jika di deklaris baris pertama aja
	*/
	Name := "Habil Arifin"
	fmt.Println(Name)

	Name = "Muhammad Habil"
	fmt.Println(Name)

	var(
		firstName = "Muhammad"
		middledName ="Habil"
		lastName = "Arifin"
	)

	fmt.Println(firstName)
	fmt.Println(middledName)
	fmt.Println(lastName)
}