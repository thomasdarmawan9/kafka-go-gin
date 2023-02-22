package food_domain

import (
	"fmt"
	"orders/db"
	"orders/utils/error_formats"
	"orders/utils/error_utils"
)

var FoodRepo foodDomain = &foodRepo{}

const (
	queryGetFoodById     = `SELECT id, name, price, stock FROM foods WHERE id = $1`
	queryCreateFood      = `INSERT INTO foods(id, name, price, stock) VALUES($1, $2, $3, $4)`
	queryReduceFoodStock = `UPDATE foods SET stock = stock - $2 WHERE id = $1`
)

type foodDomain interface {
	GetFoodById(id int32) (*Food, error_utils.MessageErr)
	CreateFood(*Food) error_utils.MessageErr
	ReduceFoodStock(*Food) error_utils.MessageErr
}

type foodRepo struct{}

func (f *foodRepo) GetFoodById(id int32) (*Food, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryGetFoodById, id)

	var food Food

	if err := row.Scan(&food.Id, &food.Name, &food.Price, &food.Stock); err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &food, nil
}

func (f *foodRepo) CreateFood(foodReq *Food) error_utils.MessageErr {
	db := db.GetDB()

	_, err := db.Exec(queryCreateFood, foodReq.Id, foodReq.Name, foodReq.Price, foodReq.Stock)

	if err != nil {
		fmt.Println(err)
		return error_formats.ParseError(err)
	}

	return nil
}

func (f *foodRepo) ReduceFoodStock(foodReq *Food) error_utils.MessageErr {
	db := db.GetDB()

	_, err := db.Exec(queryReduceFoodStock, foodReq.Id, foodReq.Stock)

	if err != nil {
		return error_formats.ParseError(err)
	}

	return nil
}
