package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukan bilangan pertama: ")
	fmt.Scanln(&a)

	fmt.Print("Masukan bilangan kedua")
	fmt.Scanln(&b)

	if a > b {
		fmt.Printf("%d lebih besar dari %d\n", a, b) // %d = integer (bilangan bulat) // \n	newline (pindah baris)
	} else if a < b {
		fmt.Printf("%d lebih besar dari %d\n", a, b)
	} else {
		fmt.Println("kedua bilangan sama")
	}
}