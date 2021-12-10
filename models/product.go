package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/shared-types"
)

type Product struct {
	ID            *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title         *string   `json:"title,omitempty" validate:"required"`
	Categories    *[]string `json:"categories,omitempty" validate:"required"`
	ImageUrl      *string   `json:"imageUrl,omitempty" validate:"required"`
	Price         *int      `json:"price,omitempty" validate:"required"`
	DiscountPrice *int      `json:"discountPrice,omitempty"`
	UnitType      *string   `json:"unitType,omitempty" validate:"required"`
	Information   *string   `json:"information,omitempty" validate:"required"`
	Nutrition     *string   `json:"nutrition,omitempty" validate:"required"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}

func (c Product) Validate(product Product) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(product)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element sharedTypes.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
