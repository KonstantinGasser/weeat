package record

type FoodCategory int

const (
	Fruit FoodCategory = iota + 1 // +1 required since database id starts at 1
	Vegetable
	Meat
	Fish
	Dairy
	Grains
)
