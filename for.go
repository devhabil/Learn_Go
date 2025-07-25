package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perlulangan ke ", counter)
	// 	counter++
	// }

	
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perlulangan ke ", counter)
		
	}
	fmt.Println("Selesai")

	names := []string{"Muhammad", "Habil", "Arifin"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names{
		fmt.Println("Index", index, "=", name)
	}
	for _, name := range names{
		fmt.Println(name)
	}
}