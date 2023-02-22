package food_domain

import (
	"time"
)

type Food struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Stock     int8      `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}
