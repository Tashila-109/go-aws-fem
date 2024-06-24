package main

import (
	"fmt"
	"go-aws-fem/imports"
)

func main() {
	newTicket := imports.Ticket{
		ID:    123,
		Event: "FEM",
	}

	newTicket.PrintEvent()

	fmt.Println(newTicket)
}
