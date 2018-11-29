package model

import "time"

type EventsUsers struct {
	Id int64
	EventsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Events *Events 
	UserId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	User *User
	Remark string
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}