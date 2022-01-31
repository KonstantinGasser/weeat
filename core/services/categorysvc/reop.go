package categorysvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/domain/category"
)

type CategoryRepo interface {
	Get(ctx context.Context, _type category.Type) ([]category.Category, error)
}
