package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (ib *NormalBuilder) setWindowType() {
	ib.windowType = "Wooden Window"
}

func (ib *NormalBuilder) setDoorType() {
	ib.doorType = "Wooden Door"
}

func (ib *NormalBuilder) setNumFloor() {
	ib.floor = 6
}

func (ib *NormalBuilder) getHouse() House {
	return House{
		doorType:   ib.doorType,
		windowType: ib.windowType,
		floor:      ib.floor,
	}
}
