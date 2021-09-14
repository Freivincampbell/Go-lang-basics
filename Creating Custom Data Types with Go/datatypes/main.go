package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Freivin", "Campbell")
	err := p.SetTwitterHandler("@freivincampbell")
	if err != nil {
		fmt.Printf("an error occurred setting twitter handeler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.ID())
	println(p.FullName())
}
