package creational

// Builder - Creational Pattern
// Separates construction of complex object from its representation

type House struct {
	windows int
	doors   int
	floors  int
	hasGarage bool
	hasPool bool
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) WithWindows(count int) *HouseBuilder {
	b.house.windows = count
	return b
}

func (b *HouseBuilder) WithDoors(count int) *HouseBuilder {
	b.house.doors = count
	return b
}

func (b *HouseBuilder) WithFloors(count int) *HouseBuilder {
	b.house.floors = count
	return b
}

func (b *HouseBuilder) WithGarage() *HouseBuilder {
	b.house.hasGarage = true
	return b
}

func (b *HouseBuilder) WithPool() *HouseBuilder {
	b.house.hasPool = true
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

// Usage: house := NewHouseBuilder().WithWindows(10).WithDoors(2).WithGarage().Build()
