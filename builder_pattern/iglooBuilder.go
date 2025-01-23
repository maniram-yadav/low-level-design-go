package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (ib *IglooBuilder) setWindowType() {
	ib.windowType = "Snow Window"
}

func (ib *IglooBuilder) setDoorType() {
	ib.doorType = "Snow Door"
}

func (ib *IglooBuilder) setNumFloor() {
	ib.floor = 4
}

func (ib *IglooBuilder) getHouse() House {
	return House{
		doorType:   ib.doorType,
		windowType: ib.windowType,
		floor:      ib.floor,
	}
}
