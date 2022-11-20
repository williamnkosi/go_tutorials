package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "mwe@gmail.com",
			zipcode: 23456,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("WillDawg")
	jim.print()

}

func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
