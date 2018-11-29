package model

import "time"

type Adventures struct {
	Id int64
	Name string `sql:",notnull"`
	Description string
	FlowsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}