package main

import "fmt"

func main() {

	type NoKTP string

	var ktpHabil NoKTP = "1111111111"

	var contoh string = "2222222222"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpHabil)
	fmt.Println(contohKtp)
}