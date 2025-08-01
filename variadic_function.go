package main

import "fmt"

func summAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(summAll(10, 10, 10))
	fmt.Println(summAll(10, 10, 10, 10, 10, 10))
	fmt.Println(summAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(summAll(numbers...))
}