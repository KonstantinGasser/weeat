package records

import (
	"context"
	"errors"

	"github.com/KonstantinGasser/weeat/core/domain/record"
	"github.com/KonstantinGasser/weeat/core/dto"
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

	if ok := newFood.Valid(); !ok {
		return errors.New("create food: missing information")
	}

	// store in db
	if err := svc.repo.InsertFood(ctx, newFood); err != nil {
		return errors.Wrap(err, "could not store food in database")
	}
	return nil
}
