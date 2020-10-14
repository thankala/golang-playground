package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var nick person
	// fmt.Println(nick)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Whatever",
		contactInfo: contactInfo{
			email:   "whatever@whatever.com",
			zipCode: 2133,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
