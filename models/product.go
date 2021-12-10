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

func (c Product) ProductValid() bool {
	if c.Title == nil {
		return false
	} else if c.Categories == nil {
		return false
	} else if c.ImageUrl == nil {
		return false
	} else if c.Price == nil {
		return false
	} else if c.UnitType == nil {
		return false
	} else if c.Information == nil {
		return false
	} else if c.Nutrition == nil {
		return false
	} else {
		return true
	}
}
