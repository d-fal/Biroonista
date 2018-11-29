package model

import "time"

const modelname = "Exchange"

type Exchange struct{
	Id int64
	Name string `sql:"unique"`
	fx float64 `sql:",notnull"`
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}
