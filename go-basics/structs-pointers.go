package main

import "fmt"

type engineer struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string
	phone string
}

func (pointerToEngineer *engineer) updateName(newFirstName string) {
	(*pointerToEngineer).firstName = newFirstName
}

func (e engineer) print() {
	fmt.Printf("%+v", e)
}

func basicStructs(){
	//erol := engineer{"Erol", "Hira"}
	erol := engineer{
		firstName: "Erol",
		lastName: "Hira",
		//contactInfo --> this is same as --> contactInfo contactInfo
		contact: contactInfo{
			email: "hira.erol@gmail.com",
			phone: "123123123",
		},
	}
	fmt.Println(erol)
	fmt.Printf("%+v", erol) //to log as formatted

	var zeynep engineer
	zeynep.firstName = "Zeynep"
	zeynep.lastName = "Hira"
	zeynep.print()

	//call over pointer --> pass by reference
	erolPointer := &erol
	erolPointer.updateName("erol")
	erol.print()

	//call over pointer -- shortcut
	/*
	If we define a receiver with a type of pointer;
	go will allow us to either call the function with a pointer or with the root type.
	go will automatically turn your variable into a pointer.
	 */
	erol.updateName("Eerol")
	erol.print()
}