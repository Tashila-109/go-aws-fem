package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(newName string) {
	p.Name = newName
}

// go chose composition over inheritance
func main() {
	myPerson := NewPerson("Tashila", 10)
	myPerson.changeName("Tashi")

	// Pointers are used to pass a reference to a value
	a := 7
	b := &a
	*b = 10
	fmt.Println(b)
	fmt.Println(a)

	fmt.Println(myPerson)

	mySlice := []int{1, 2, 3, 4, 5}

	for index := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}
