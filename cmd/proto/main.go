package main

import (
	"fmt"

	addressbookpb "github.com/jeffpignataro/golang/cmd/proto/models/addressbooks/addressbookpb"
	personpb "github.com/jeffpignataro/golang/cmd/proto/models/persons/personpb"
)

func main() {
	var a addressbookpb.AddressBook
	var p *personpb.Person
	p.Name = "Jeff Pignataro"
	a.People[0] = p
	fmt.Println(a.People[0])
}
