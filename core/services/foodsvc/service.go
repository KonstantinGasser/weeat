package foodsvc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/domain/food"
	"github.com/KonstantinGasser/weeat/core/dto"
	"github.com/KonstantinGasser/weeat/pkg/http/response"
	"github.com/sirupsen/logrus"
)

const (
	searchLimit = 10
)

type Service struct {
	repo RecordsRepo
}

func New(repo RecordsRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc Service) InsertFood(ctx context.Context, foodItem dto.Food) response.RespErr {

	newFood := food.FoodFromDTO(foodItem)

	if err := newFood.Valid(); err != nil {
		if err == food.ErrSugarBiggerCarbs {
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

	foodItem := food.FoodFromDAO(item).Scale(amount)

	return dto.Food{
		ID:       foodItem.ID,
		Name:     foodItem.Name,
		Category: int(foodItem.Category),
		Kcal:     foodItem.Kcal.Value(),
		Carbs:    foodItem.Carbs.Value(),
		Sugar:    foodItem.Sugar.Value(),
		Protein:  foodItem.Protein.Value(),
		Fats:     foodItem.Fats.Value(),
	}, nil
}

func (svc Service) SearchFood(ctx context.Context, query string, limit int) ([]dto.FoodQuery, response.RespErr) {

	// limit is the limit for the SQL-Statement, it should not be to big
	if limit > searchLimit {
		return nil, response.Err(fmt.Errorf("limit exceeds allows threshold"), http.StatusBadRequest, "Search limit must be lower 10")
	}

	queryItems, err := svc.repo.SearchFood(ctx, query, limit)
	if err != nil {
		logrus.Errorf("test %v\n", err.Error())
		return nil, response.Err(err, http.StatusInternalServerError, "could not lookup query")
	}

	return dto.QueryFromDAO(queryItems), nil
}
