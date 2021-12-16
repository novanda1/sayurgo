package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/shared-types"
)

// admin struct to describe admin object
type Admin struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName *string   `json:"displayName,omitempty" validate:"required"`
	Phone       *string   `json:"phone,omitempty" validate:"required"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

func (c Admin) Validate(admin Admin) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(admin)
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
