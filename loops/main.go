package main

import (
	"fmt"
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "bird")

	for index, value := range animals {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	for value := range 10 {
		fmt.Println(value)
	}

	// this is how you do while loops in go
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
