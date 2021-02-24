package model

import "github.com/go-openapi/strfmt"

// Account model.
//
//swagger:model account
type Account struct {
	// The ID of the Account
	//
	// required: true
	ID int `json:"id"`

	// Status of the order.
	//
	// required: true
	Status string `json:"status"`

	// Time the Account was created
	//
	// required: true
	CreatedAt strfmt.DateTime `json:"created_at"`

	// Time the Account was updated
	//
	// required: true
	UpdatedAt strfmt.DateTime `json:"updated_at"`

	// Time the Account was deleted
	//
	// required: true
	DeletedAt strfmt.DateTime `json:"deleted_at"`
}
