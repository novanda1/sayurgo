package models

import (
	"time"
)

type Product struct {
	ID            *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title         *string   `json:"title,omitempty"`
	Categories    *[]string `json:"categories,omitempty"`
	ImageUrl      *string   `json:"imageUrl,omitempty"`
	Price         *int      `json:"price,omitempty"`
	DiscountPrice *int      `json:"discountPrice,omitempty"`
	UnitType      *string   `json:"unitType,omitempty"`
	Information   *string   `json:"information,omitempty"`
	Nutrition     *string   `json:"nutrition,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}
