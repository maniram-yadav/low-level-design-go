package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)

	house := director.buildHouse()
	fmt.Printf("\nNormal House Door Type %s ", house.doorType)
	fmt.Printf("\nNormal House Window Type %s ", house.windowType)
	fmt.Printf("\nNormal House Floor count %d", house.floor)

	director.setBuilder(iglooBuilder)
	house = director.buildHouse()

	fmt.Printf("\nNormal House Door Type %s ", house.doorType)
	fmt.Printf("\nNormal House Window Type %s ", house.windowType)
	fmt.Printf("\nNormal House Floor count %d ", house.floor)

}
