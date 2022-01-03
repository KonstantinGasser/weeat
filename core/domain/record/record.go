package record

type Food struct {
	Name     string
	Category FoodCategory
	Kcal     int
	Carbs    float64
	Protein  float64
	Fats     float64
}

type Recipe struct {
	Name        string
	Ingredients []Food
}
