package model

import "time"

type Membership struct{
	Id int64
	Name string
	Price int64
	CurrencyId int64
	RolesList string
	ExchangeId int64
	Exchange *Exchange
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}

