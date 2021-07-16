package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of person which is a slice of strings
type people []string

//A receiver sets up methods on variables that we create.
//receiver p --> its type is person
//any variable of type people now gets access to the print method.
//by convention we usually refer to the receiver with a one or two letter abbreviation
func (p people ) print()  {
	for i, person := range p {
		fmt.Println(i, ":", person)
	}
}

func createPeople() people {
	peopleVar := people{}

	names := []string{"John", "Mohn", "Than", "Khan"}
	surnames := []string{"Lahn", "Mahn"}

	for _, name := range names {
		for _, surname := range surnames {
			peopleVar = append(peopleVar, name + " " + surname)
		}
	}

	return peopleVar
}

/**
function that returns multiple values
 */
func read(p people, count int) (people, people) {
	return p[:count], p[count:]
}

//joining a slice of strings
func (p people) toString() string {
	return strings.Join(p, ",")
}

func (p people) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, convertToBytes(p.toString()), 0666)
}

func readFromFile(fileName string) people {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return people(s)
}

func (p people) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range p {
		//generate a number between 0 and the length-1 of the slice
		newPosition := r.Intn(len(p) - 1)
		p[i], p[newPosition] = p[newPosition], p[i]
	}
}
