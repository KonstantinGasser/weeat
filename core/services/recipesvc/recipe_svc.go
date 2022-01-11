package recipesvc

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
)

type Service struct {
	repo interface{}
}

func (svc Service) InsertRecipe(ctx context.Context, recipe dto.Recipe) response.RespErr {

}
