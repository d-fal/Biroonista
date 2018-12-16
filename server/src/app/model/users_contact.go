package model

import (
	"time"
)

type UsersContacts struct {
	Model
	UsersId          int64
	Users            *Users
	Address          string
	IsMain           bool
	IsVerified       bool
	ContactsType     int64
	Verifier         string
	VerificationDate time.Time
	Deleatables
}

func (u *UsersContacts) SetUserContacts(contact string) {
	u.Address = contact
}
