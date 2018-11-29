package model

import "time"

type Services struct {
	Id int64
	Name string
	Description string
	Cost float64
	ExchangeId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Exchange *Exchange
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}