package records

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
)

type RecordsRepo interface {
	InsertFood(ctx context.Context, food dao.Food) error
	SearchFood(ctx context.Context, query string) ([]dao.FoodQuery, error)
	UpdateFood(ctx context.Context, column string, value interface{}) error
	DeleteFood(ctx context.Context, ID int) error

	// InsertRecipe(ctx context.Context)
	// UpdateRecipe(ctx context.Context)
	// DeleteRecipe(ctx context.Context)
}
