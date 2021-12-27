package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/pkg/shared-types"
)

var UserCollectionName = "users"

type UserRole string

var (
	Customer UserRole = "customer"
	Admin    UserRole = "Admin"
)

func (s UserRole) String() string {
	switch s {
	case Admin:
		return "admin"
	}
	return "customer"
}

type UserAddress struct {
	ID         *string `json:"id,omitempty" bson:"_id,omitempty"`
	Title      *string `json:"title,omitempty" validate:"required"`
	Recipient  *string `json:"recipient,omitempty" validate:"required"`
	Phone      *string `json:"phone,omitempty" validate:"required"`
	City       *string `json:"city,omitempty" validate:"required"`
	PostalCode *string `json:"postalCode,omitempty" validate:"required"`
	Address    *string `json:"address,omitempty" validate:"required"`
	Detail     *string `json:"detai,omitempty"`
}

type User struct {
	ID          *string        `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName *string        `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Phone       *string        `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	UserAddress *[]UserAddress `json:"userAddress,omitempty" bson:"userAddress,omitempty"`
	Role        *string        `json:"role,omitempty" bson:"role,omitempty"`
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

func (c UserAddress) Validate(address UserAddress) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(address)
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
