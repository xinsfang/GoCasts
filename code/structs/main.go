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
    //alex := person{"Alex", "Anderson", contactInfo{"abc@def.com", 123456}}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	jim.updateName("jimmy") // go allows us to use a pointer shortcut

	jim.print()
}

// go is a pass-by-value language, this faulty func is actually updating a copy of original person
//func (person person) updateName(newFirstName string) {
//	person.firstName = newFirstName
//}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	//fmt.Printf("%v", p) //the value in a default format
	fmt.Printf("%+v", p) //adds field names
}
