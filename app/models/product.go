package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/pkg/shared-types"
)

var ProductCollectionName = "products"

type Product struct {
	ID            *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title         *string   `json:"title,omitempty" validate:"required"`
	Categories    *[]string `json:"categories,omitempty" validate:"required"`
	ImageUrl      *string   `json:"imageUrl,omitempty" validate:"required,url"`
	Price         *int      `json:"price,omitempty" validate:"required,numeric"`
	DiscountPrice *int      `json:"discountPrice,omitempty"`
	UnitType      *string   `json:"unitType,omitempty" validate:"required"`
	Information   *string   `json:"information,omitempty" validate:"required,alphanum"`
	Nutrition     *string   `json:"nutrition,omitempty" validate:"required,alphanum"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}

type GetAllProductsParams struct {
	Limit int64 `json:"limit" validate:"required,numeric"`
	Page  int64 `json:"page" validate:"required,numeric"`
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

func (c GetAllProductsParams) Validate(params GetAllProductsParams) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(params)
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
