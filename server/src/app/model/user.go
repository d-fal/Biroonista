package model

import "github.com/paulmach/orb"

type User struct {
	Model
	FirstName string
	MiddleName string
	LastName string
	UserName string `sql:",notnull, unique"`
	Password string `sql:",notnull"`
	Loc orb.Point `sql:",type:Point"`
	TimeZone string
	IsVerified bool
	isActive bool
	MembershipId int64
	Contacts []Contacts `pg:"many2many:user_to_contacts"`
	Bank float64
	ExchangeId int64
	Exchange *Exchange
	Rating int `sql:"-"`
	Deleatables
}

func (u *User) GetId() int64 {
	return u.Id
}

