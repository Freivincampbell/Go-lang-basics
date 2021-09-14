package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Freivin", "Campbell")
	err := p.SetTwitterHandler("@freivincampbell")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("an error occurred setting twitter handeler: %s\n", err.Error())
	}

	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.FullName())
}
