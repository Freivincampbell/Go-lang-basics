package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	handle string
	name   string
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")

	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "Costa Rica"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	name Name
}

type Person struct {
	Name
	first          string
	last           string
	twitterHandler TwitterHandler
	Citizen
}

func NewPerson(fistName, lastName string, citizen Citizen) Person {
	return Person{
		Citizen: citizen,
		Name: Name{
			first: fistName,
			last:  lastName,
		},
	}
}

func (p *Person) ID() string {
	return fmt.Sprintf("person's identifier: %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
