package main

import "fmt"

func getComplateName() (firstName, middleName, LastName string) {
	firstName = "Muhammad"
	middleName = "Habil"
	LastName = "Arifin"

	return firstName, middleName, LastName
}

func main() {
	a, b, c := getComplateName()
	fmt.Println(a,b,c)

}