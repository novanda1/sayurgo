package models

import (
	"time"
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
	Phone       *string        `json:"phone,omitempty"`
	UserAddress *[]UserAddress `json:"userAddress,omitempty"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
}
