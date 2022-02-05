package verify

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
)

type VerifyRepo interface {
	InitVerifyItem(ctx context.Context, item dao.VerifyItem) error
}
