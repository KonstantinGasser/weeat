package records

import (
	"context"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/domain/record"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/pkg/errors"
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

func (svc Service) InsertFood(ctx context.Context, food dto.Food) error {

	newFood := record.FoodFromDTO(food)

	// if ok := newFood.Valid(); !ok {
	// 	return errors.New("create food: missing information")
	// }

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
		return errors.Wrap(err, "recordsvc.InsertFood storing item")
	}
	return nil
}
