package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	sharedTypes "github.com/novanda1/sayurgo/pkg/shared-types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var OrderCollectionName = "orders"

type GetAllOrdersParams struct {
	Limit int64 `json:"limit" validate:"required,numeric"`
	Page  int64 `json:"page" validate:"required,numeric"`
}

type OrderStatus string

var (
	Waiting    OrderStatus = "waiting"
	Process    OrderStatus = "process"
	OnDelivery OrderStatus = "on delivery"
	Completed  OrderStatus = "completed"
	Issued     OrderStatus = "issued"
)

func (s OrderStatus) String() string {
	switch s {
	case Waiting:
		return "waiting"
	case Process:
		return "process"
	case OnDelivery:
		return "on delivery"
	case Completed:
		return "completed"
	case Issued:
		return "issued"
	}

	return "waiting"
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
	Status     string             `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty"`
}

type OrderResponseData struct {
	HasNext bool     `json:"hasNext"`
	Orders  []*Order `json:"orders"`
}

type OrderResponse struct {
	Success bool              `json:"success"`
	Data    OrderResponseData `json:"data"`
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

	return errors
}
