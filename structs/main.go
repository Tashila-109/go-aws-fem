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

	fmt.Println(myPerson)
}
