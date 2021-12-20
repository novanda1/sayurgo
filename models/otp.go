package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/shared-types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// would be much better if we use redis.
type Otp struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Phone *string            `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Otp   *string            `json:"otp,omitempty" bson:"otp,omitempty"`
	Exp   time.Time          `json:"exp,omitempty"`
}

type VerifOtpResult struct {
	Verified *bool   `json:"verified"`
	User     *User   `json:"user"`
	Token    *string `json:"token"`
}

func (c Otp) Validate(otp Otp) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(otp)
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
