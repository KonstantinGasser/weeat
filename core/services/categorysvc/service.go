package categorysvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/domain/category"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
)

type Serivce struct {
	repo CategoryRepo
}

func New(repo CategoryRepo) *Serivce {
	return &Serivce{
		repo: repo,
	}
}

func (svc Serivce) Get(ctx context.Context, _type int) ([]category.Category, response.RespErr) {
	return nil, nil
}
