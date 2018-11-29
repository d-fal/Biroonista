package model

import "time"

type Attachments struct {
	Id int64
	Name string `sql:",notnull"`
	Description string
	TypeCode int `sql:",notnull"`
	AttachmentsId int64
	Attachments *Attachments
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}