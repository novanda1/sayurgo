package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/shared-types"
)

type UserAddress struct {
	ID         *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title      *string   `json:"title,omitempty"`
	Recipient  *string   `json:"recipient,omitempty"`
	Phone      *string   `json:"phone,omitempty"`
	City       *string   `json:"city,omitempty"`
	PostalCode *string   `json:"postalCode,omitempty"`
	Address    *string   `json:"address,omitempty"`
	Detail     *string   `json:"detai,omitemptyl"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

type User struct {
	ID          *string        `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName *string        `json:"displayName,omitempty"`
	Phone       *string        `json:"phone,omitempty" validate:"required"`
	UserAddress *[]UserAddress `json:"userAddress,omitempty"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
}

func (c User) AuthDtoValidate(user User) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(user)
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
