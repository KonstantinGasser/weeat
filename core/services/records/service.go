package records

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/domain/record"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
)

// 1. Insert/Update/Delete:
// 		-> Foods with neutritions
// 		-> recepies with neutritions

type Service struct {
	repo RecordsRepo
}

func New(repo RecordsRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc Service) InsertFood(ctx context.Context, food dto.Food) response.RespErr {

	newFood := record.FoodFromDTO(food)

	if err := newFood.Valid(); err != nil {
		if err == record.ErrSugarBiggerCarbs {
			return response.Err(err, http.StatusBadRequest, "Sugar Value must be lower then or equal to carbs value")
		}
		return response.Err(err, http.StatusBadRequest, "Provided data is invalid")
	}

	// store in db
	daoFood := dao.Food{
		Name:     newFood.Name,
		Category: int(newFood.Category),
		Kcal:     newFood.Kcal,
		Carbs:    newFood.Carbs,
		Protein:  newFood.Protein,
		Fats:     newFood.Fats,
	}
	if err := svc.repo.InsertFood(ctx, daoFood); err != nil {
		return response.Err(err, http.StatusInternalServerError, "Could not save food item")
	}
	return nil
}
