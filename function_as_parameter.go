package main

import "fmt"

// program function type declaration
type Filter func(string) string


func sayHelloWithFillter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)

// func sayHelloWithFillter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string{
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFillter("Habil", spamFilter)

	filter := spamFilter
	sayHelloWithFillter("Anjing", filter)
}