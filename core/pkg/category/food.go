package category

type Type int

const (
	Fruit Type = iota + 1 // +1 required since database id starts at 1
	Vegetable
	Meat
	Fish
	Dairy
	Grains
	SoftDrinks
	Alcohol
	BreakFast
	Lunch
	Dinner
	Snack
)
