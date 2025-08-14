package main

import "fmt"

type Hasname interface {
	GetName() string
}

func sayHello(value Hasname) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return  animal.Name
}

func main() {
	person := Person{Name: "Habil"}
	sayHello(person)

	animal := Animal{Name: "Kucing"}
	sayHello(animal)
}