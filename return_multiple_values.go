package main

import "fmt"

func getFulName() (string, string, string) {
	return "Muhammad", "Habil", "Arifin"
}

func main() {
	// firstName, middleName, LastName := getFulName()
	// fmt.Println(firstName, middleName, LastName)
	
	firstName, _, _:= getFulName()
	fmt.Println(firstName)
}