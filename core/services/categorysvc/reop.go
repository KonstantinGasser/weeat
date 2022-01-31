package categorysvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
)

type CategoryRepo interface {
	GetCategories(ctx context.Context, _type int) ([]dao.Category, error)
}
