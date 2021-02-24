package model

/*
Table Account{
    id int [pk, increment]
    status varchar
    created_at datetime
    updated_at datetime
    deleted_at datetime
  }
*/
type Account struct {
	id     int
	status string

	createdAt string
	updatedAt string
	deletedAt string
}
