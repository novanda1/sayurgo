package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/shared-types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus struct{ Status string }

var (
	Waiting    = OrderStatus{"waiting"}
	Process    = OrderStatus{"process"}
	OnDelivery = OrderStatus{"on delivery"}
	Completed  = OrderStatus{"completed"}
	Issued     = OrderStatus{"issued"}
)

func (r OrderStatus) String() string {
	return r.Status
}

type OrderProduct struct {
	ProductID    primitive.ObjectID `json:"productID,omitempty" bson:"productID,omitempty" validate:"required"`
	TotalProduct *int               `json:"totalProduct,omitempty" bson:"totalProduct,omitempty" validate:"required"`
	AtPrice      *int               `json:"atPrice,omitempty" bson:"atPrice,omitempty" validate:"required"`
}

type Order struct {
	ID         *string            `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userid,omitempty" bson:"userid,omitempty"`
	Products   *[]OrderProduct    `json:"products,omitempty" bson:"products,omitempty"`
	TotalPrice *int               `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	Status     *OrderStatus       `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty"`
}

func (c Order) Validate(order Order) []*sharedTypes.ErrorResponse {
	var errors []*sharedTypes.ErrorResponse
	validate := validator.New()

	err := validate.Struct(order)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element sharedTypes.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	fmt.Println(OrderStatus.String(*order.Status))

	return errors
}
