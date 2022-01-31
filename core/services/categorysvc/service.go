package categorysvc

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dto"
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

func (svc Serivce) Get(ctx context.Context, _type int) ([]dto.Category, response.RespErr) {

	cats, err := svc.repo.GetCategories(ctx, _type)
	if err != nil {
		return nil, response.Err(err, http.StatusInternalServerError, "could not get categories")
	}

	var categories []dto.Category
	for _, c := range cats {
		categories = append(categories, dto.Category{
			ID:    c.ID,
			Type:  int(c.Type),
			Label: c.Label,
			Emoji: c.Emoji,
		})
	}
	return categories, nil
}
