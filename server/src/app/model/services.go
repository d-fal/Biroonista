package model

type Services struct {
	Model
	Name string
	Description string
	Cost float64
	ExchangeId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Exchange *Exchange
	Deleatables
}