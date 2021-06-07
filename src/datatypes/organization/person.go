package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialsecurityNumber string

func NewSocialSecurityNumber(value string) Citizen{
    return socialsecurityNumber(value)
}

func (ssn socialsecurityNumber) ID() string {
    return string(ssn)
}

func (ssn socialsecurityNumber) Country() string {
	return "The United States"
}

type europeanUnionIdentifier struct {
	id string
	country string
}

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
        case string:
	        return europeanUnionIdentifier{
	            id: v,
	            country: country,
	        }
	    case int:
	    	return europeanUnionIdentifier{
                id: strconv.Itoa(v),
                country: country,
            }
	    default:
		    panic("using an invalid type to initialize eu")
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
    return fmt.Sprintf("eu: %s", eui.country)
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
    cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Name struct {
	first string
	last string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
    twitterHandler TwitterHandler
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
    return Person{
    	Name: Name{
    		first: firstName,
    	    last: lastName,
    	},
    	Citizen: citizen,
    }
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier is %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
    if len(handler) == 0 {
    	p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with @")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}