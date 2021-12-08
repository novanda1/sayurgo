package models

import (
	"time"
)

type CartProduct struct {
	ID           *string `json:"id,omitempty" bson:"_id,omitempty"`
	TotalProduct *int    `json:"totalProduct"`
	ProductID    *string `json:"productID"`
}

type Cart struct {
	ID         *string        `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     *string        `json:"userid"`
	TotalPrice *int           `json:"totalPrice"`
	Product    *[]CartProduct `json:"product"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}
