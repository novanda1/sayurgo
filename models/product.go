package models

import (
	"time"
)

type Product struct {
	ID            *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title         *string   `json:"title"`
	Categories    *[]string `json:"categories"`
	ImageUrl      *string   `json:"imageUrl"`
	Price         *int      `json:"price"`
	DiscountPrice *int      `json:"discountPrice"`
	UnitType      *string   `json:"unitType"`
	Information   *string   `json:"information"`
	Nutrition     *string   `json:"nutrition"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
