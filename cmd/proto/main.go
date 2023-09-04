package main

import (
	"fmt"

	p "github.com/jeffpignataro/golang/cmd/proto/models/persons/person"
)

func main() {
	// var a addressbook.AddressBook
	p := p.Person{
		Name:  "My Name",
		Id:    0,
		Email: "my@email.com",
		Phones: []*p.Person_PhoneNumber{
			{
				Number: "123-456-7890",
				Type:   p.Person_HOME,
			},
		},
	}
	fmt.Println(p)
}
