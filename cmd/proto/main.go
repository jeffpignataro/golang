package main

import (
	"fmt"
	"workspace/bazel-bin/cmd/proto/models/persons/person_go_proto_/github.com/jeffpignataro/golang/cmd/proto/models/persons/person"
)

func main() {
	// var a addressbook.AddressBook
	p := person.Person{
		Name:  "My Name",
		Id:    0,
		Email: "my@email.com",
		Phones: []*person.Person_PhoneNumber{
			&person.Person_PhoneNumber{
				Number: "123-456-7890",
				Type:   person.Person_HOME,
			},
		},
	}
	fmt.Println(p)
}
