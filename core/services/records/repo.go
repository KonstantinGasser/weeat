package records

import (
	"context"

	"github.com/KonstantinGasser/lowco/core/dao"
)

type RecordsRepo interface {
	InsertFood(ctx context.Context, food dao.Food) error
	UpdateFood(ctx context.Context, column string, value interface{}) error
	DeleteFood(ctx context.Context, ID int) error

	// InsertRecipe(ctx context.Context)
	// UpdateRecipe(ctx context.Context)
	// DeleteRecipe(ctx context.Context)
}
