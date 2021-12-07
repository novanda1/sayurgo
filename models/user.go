package models

import (
	"time"
)

type UserAddress struct {
	ID         *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title      *string   `json:"title"`
	Recipient  *string   `json:"recipient"`
	Phone      *string   `json:"phone"`
	City       *string   `json:"city"`
	PostalCode *string   `json:"postalCode"`
	Address    *string   `json:"address"`
	Detail     *string   `json:"detail"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type User struct {
	ID          *string        `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName *string        `json:"displayName"`
	Phone       *string        `json:"phone"`
	UserAddress *[]UserAddress `json:"userAddress"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}
