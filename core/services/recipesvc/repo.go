package recipesvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
)

type RecordsRepo interface {
	InsertFood(ctx context.Context, food dao.Food) error
	GetFood(ctx context.Context, ID string) (dao.Food, error)
	SearchFood(ctx context.Context, query string, limit int) ([]dao.FoodQuery, error)
	UpdateFood(ctx context.Context, column string, value interface{}) error
	DeleteFood(ctx context.Context, ID int) error

	InsertRecipe(ctx context.Context, recipe dao.Recipe) (int, error)
}
