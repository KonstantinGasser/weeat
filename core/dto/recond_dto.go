package dto

import "github.com/KonstantinGasser/weeat/core/dao"

type Food struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category int     `json:"food_cat"`
	Kcal     float64 `json:"kcal"`
	Carbs    float64 `json:"carbs"`
	Sugar    float64 `json:"sugar"`
	Protein  float64 `json:"protein"`
	Fats     float64 `json:"fats"`
}

type FoodQuery struct {
	ID       int
	Name     string
	Category int
}

func QueryFromDAO(items []dao.FoodQuery) []FoodQuery {
	var res []FoodQuery

	for _, item := range items {
		res = append(res, FoodQuery{
			ID:       item.ID,
			Name:     item.Name,
			Category: item.Category,
		})
	}
	return res
}
