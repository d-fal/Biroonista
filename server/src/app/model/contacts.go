package model

import "time"
type Contacts struct {
	Id int64
	Address string
	IsMain bool
	IsVerified bool
	ContactstypeId int64
	Contactstype *Contactstype
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}
