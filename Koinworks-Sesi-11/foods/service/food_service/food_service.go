package food_service

import (
	"foods/domain/food_domain"
	"foods/utils/error_utils"
)

var FoodService foodServiceInterface = &foodService{}

type foodServiceInterface interface {
	GetFoodById(id int32) (*food_domain.Food, error_utils.MessageErr)
	CreateFood(*food_domain.Food) (*food_domain.Food, error_utils.MessageErr)
	ReduceFoodStock(*food_domain.Food) error_utils.MessageErr
}

type foodService struct{}

func (f *foodService) GetFoodById(id int32) (*food_domain.Food, error_utils.MessageErr) {
	food, err := food_domain.FoodRepo.GetFoodById(id)

	if err != nil {
		return nil, err
	}

	return food, nil
}

func (f *foodService) CreateFood(foodReq *food_domain.Food) (*food_domain.Food, error_utils.MessageErr) {

	err := foodReq.Validate()

	if err != nil {
		return nil, err
	}

	food, err := food_domain.FoodRepo.CreateFood(foodReq)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (f *foodService) ReduceFoodStock(foodReq *food_domain.Food) error_utils.MessageErr {
	err := food_domain.FoodRepo.ReduceFoodStock(foodReq)

	if err != nil {
		return err
	}

	return nil
}
