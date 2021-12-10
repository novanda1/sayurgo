package models

import (
	"time"
)

type CartProduct struct {
	ID           *string `json:"id,omitempty" bson:"_id,omitempty"`
	TotalProduct *int    `json:"totalProduct,omitempty" bson:"totalProduct,omitempty" binding:"required"`
	ProductID    *string `json:"productID,omitempty" bson:"productID,omitempty" binding:"required"`
}

type Cart struct {
	ID         *string        `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     *string        `json:"userid,omitempty" bson:"userid,omitempty" binding:"required"`
	TotalPrice *int           `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	Product    *[]CartProduct `json:"product,omitempty" bson:"product,omitempty"`
	CreatedAt  time.Time      `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time      `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
