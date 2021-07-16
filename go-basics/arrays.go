package main

import "fmt"

/**
- slice is an array that can grow or shrink
- every element in an array or slice must be of the same type
- fruits = []string{ "apple", "orange", "banana",  "grape", "melon" }
fruits[startIndexIncluding:upToExcluding]
fruits[0:2] --> { "apple", "orange" }
fruits[:2] --> from beginning to 2 --> { "apple", "orange" }
fruits[2:] --> from 2 to the end --> { "banana",  "grape", "melon" }
*/
func arrays(){

	sliceA := []string{ "a", "b", "c" }
	sliceA = append(sliceA, "d", "e")
	fmt.Println(sliceA)

	for i, element := range sliceA {
		fmt.Println(i, ":", element)
	}

	peopleArray := people{"John", "Mohn"}
	peopleArray.print()

	people2 := createPeople()
	people2.print()
}

func multipleReturn() {

	fmt.Println("multipleReturn-----------------")

	people := createPeople()
	fmt.Println(people.toString())

	nextPeople, remainingPeople := read(people, 2)

	nextPeople.print()
	remainingPeople.print()
}
