package category

type Food int

const (
	Fruit Food = iota + 1 // +1 required since database id starts at 1
	Vegetable
	Meat
	Fish
	Dairy
	Grains
	SoftDrinks
	Alcohol
)
