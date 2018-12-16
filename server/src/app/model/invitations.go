package model


type Invitations struct {
	Model
	name string
	price float64
	ExchangeId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Exchange *Exchange
	Deleatables
}