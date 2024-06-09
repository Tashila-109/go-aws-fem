package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "bird")

	animals = slices.Delete(animals, 0, 1)

	fmt.Println(animals)
}
