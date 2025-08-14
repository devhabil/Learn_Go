package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello (name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var habil Customer
	habil.Name = "Muhammad Habil Arifin"
	habil.Address = "Indonesia"
	habil.Age = 25

	fmt.Println(habil)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
}