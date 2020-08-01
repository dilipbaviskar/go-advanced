// Person struct with methods of value receiver

package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

//A person method with value receiver
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method with value receiver
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

func main() {
	var p Person
	p.FirstName = "Rob"
	p.LastName = "Pike"
	p.Dob = time.Date(1957, time.February, 17, 0, 0, 0, 0, time.UTC)
	p.Email = "pike@email.com"
	p.Location = "California"

	// Struct literal
	p1 := Person{
		FirstName: "Shiju",
		LastName:  "Varghese",
		Dob:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		Email:     "shiju@email.com",
		Location:  "Kochi",
	}
	p.PrintName()
	p.PrintDetails()
	p1.PrintName()
	p1.PrintDetails()

}
