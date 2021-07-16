//main package is executable package and must have a func named main.
package main

//import some code from another package named fmt.
//fmt is a standard library package contained in golang by default.
import (
	"fmt"
	"strconv"
)

/**
~/dev/ws2/ws_go/go-basics$ go run main.go
the main function runs automatically when the program is executed.
 */
func main() {
	//variables()
	//arrays()
	//multipleReturn()
	//savePeople()

	//fmt.Println("save and read from file ------------")
	//people := readFromFile("sample")
	//people.print()

	//shuffle()
	//basicStructs()
	//testMaps()
	//testInterfaces()
	//makeHttpCall()
	concurrency()
}

func variables() {

	//var name string = "Erol Hira"
	name := getName()
	var fullName bool = true
	var age int = 38
	var score float64 = 12.2121

	fmt.Println("name: " + name +
				", fullName: " + strconv.FormatBool(fullName) +
				", age: " + strconv.Itoa(age) +
				", score: " + fmt.Sprintf("%f", score))
}

func getName() string {
	return "Erol Hira"
}

func savePeople() {

	people := createPeople()
	people.saveToFile("sample")
}

func shuffle() {

	peopleList := createPeople()
	peopleList.shuffle()
	peopleList.print()
}