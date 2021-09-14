package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Freivin", "Campbell", organization.NewSocialSecurityNumber("CR-123-1234-12345"))
	err := p.SetTwitterHandler("@freivincampbell")

	fmt.Printf("%T\n", organization.TwitterHandler("test"))

	if err != nil {
		fmt.Printf("an error occurred setting twitter handeler: %s\n", err.Error())
	}

	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Name.FullName())
}
