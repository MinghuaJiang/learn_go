package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
    p := organization.NewPerson("James", "Wilson", organization.NewSocialSecurityNumber("123-45-6789"))
    fmt.Println(p.ID())
	fmt.Println(p.FullName())
    err := p.SetTwitterHandler("@janewil")
    if err != nil {
    	fmt.Printf("error occurs set twitter handler %s\n", err.Error())
	}

	fmt.Println(p.TwitterHandler())

	fmt.Println(p.TwitterHandler().RedirectUrl())

	p2 := organization.NewPerson(
		"James",
		"Ben",
		organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	fmt.Println(p2.ID())
	fmt.Println(p2.FullName())
	fmt.Println(p2.Country())
	p3 := organization.NewPerson(
		"Mike",
		"James",
		organization.NewEuropeanUnionIdentifier(1234, "France"))
	fmt.Println(p3.ID())
	fmt.Println(p3.FullName())
	fmt.Println(p3.Country())
}