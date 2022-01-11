package records

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/domain/record"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
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
		Label:    newFood.Label,
		Category: int(newFood.Category),
		Kcal:     newFood.Kcal,
		Carbs:    newFood.Carbs,
		Sugar:    newFood.Sugar,
		Protein:  newFood.Protein,
		Fats:     newFood.Fats,
	}
	if err := svc.repo.InsertFood(ctx, daoFood); err != nil {
		return response.Err(err, http.StatusInternalServerError, "Could not save food item")
	}
	return nil
}

func (svc Service) GetFood(ctx context.Context, ID string, amount int) (dto.Food, response.RespErr) {

	item, err := svc.repo.GetFood(ctx, ID)
	if err != nil {
		return dto.Food{}, response.Err(err, http.StatusInternalServerError, "could not lookup food")
	}

	fmt.Println("pg: ", item)
	food := record.FoodFromDAO(item).Scale(amount)
	fmt.Println("domain: ", food)
	return dto.Food{
		ID:       food.ID,
		Name:     food.Name,
		Category: int(food.Category),
		Kcal:     food.Kcal.Value(),
		Carbs:    food.Carbs.Value(),
		Sugar:    food.Sugar.Value(),
		Protein:  food.Protein.Value(),
		Fats:     food.Fats.Value(),
	}, nil
}

func (svc Service) SearchFood(ctx context.Context, query string) ([]dto.FoodQuery, response.RespErr) {

	queryItems, err := svc.repo.SearchFood(ctx, query)
	if err != nil {
		logrus.Errorf("test %v\n", err.Error())
		return nil, response.Err(err, http.StatusInternalServerError, "could not lookup query")
	}

	return dto.QueryFromDAO(queryItems), nil
}
