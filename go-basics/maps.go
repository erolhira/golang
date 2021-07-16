package main

import "fmt"

func testMaps() {
	colors := createMap()
	printMap(colors)
}

func printMap(c map[string]string) {

	for key, value := range c {
		fmt.Println(key, value)
	}
}

func createMap() map[string]string {

	//DECLARATION TYPE - 1
	//map[key]value
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4b0000",
	}
	fmt.Println(colors)

	//DECLARATION TYPE - 2
	var colors2 map[string]string //go intializes the map with zero value
	fmt.Println(colors2)

	//DECLARATION TYPE - 3
	colors3 := make(map[string]string)
	fmt.Println(colors3)

	//setting values
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	//deleting values
	delete(colors, "white")
	fmt.Println(colors)

	return colors
}