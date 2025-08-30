package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	habil := Man{"Habil"}
	habil.Married()

	fmt.Println(habil.Name)
}