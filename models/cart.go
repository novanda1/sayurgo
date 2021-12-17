package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/shared-types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartProduct struct {
	ID           *string            `json:"id,omitempty" bson:"_id,omitempty"`
	TotalProduct *int               `json:"totalProduct,omitempty" bson:"totalProduct,omitempty" validate:"required"`
	ProductID    primitive.ObjectID `json:"productID,omitempty" bson:"productID,omitempty" validate:"required"`
}

type Cart struct {
	ID         *string            `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userid,omitempty" bson:"userid,omitempty" validate:"required"`
	TotalPrice *int               `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	Product    *[]CartProduct     `json:"product,omitempty" bson:"product,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type UpdateAmountCartProductParam struct {
	TotalProduct *int `json:"totalProduct,omitempty" bson:"totalProduct,omitempty" validate:"required"`
}

func (c CartProduct) Validate(cartProduct CartProduct) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(cartProduct)
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
