package model

import "github.com/go-openapi/strfmt"

// Account model.
//
//swagger:model account
type Account struct {
	// The ID of the Account
	//
	// required: true
	// x-order: 0
	ID int `json:"id"`

	// Status of the order.
	//
	// required: true
	// x-order: 1
	Status string `json:"status"`

	// Time the Account was created
	//
	// required: true
	// x-order: 2
	CreatedAt strfmt.DateTime `json:"created_at"`

	// Time the Account was updated
	//
	// required: true
	// x-order: 3
	UpdatedAt strfmt.DateTime `json:"updated_at"`

	// Time the Account was deleted
	//
	// required: true
	// x-order: 4
	DeletedAt strfmt.DateTime `json:"deleted_at"`
}
