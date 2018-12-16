package model

type Membership struct{
	Model
	Name string
	Price int64
	CurrencyId int64
	RolesList string
	ExchangeId int64
	Exchange *Exchange
	Deleatables
}

