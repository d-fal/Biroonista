package model

import "time"

type InvitationAux struct {
	Id int64
	PhysicalPath string
	AttachmentsId int64 `pg:"on_delete:RESTRICT"`
	Attachments *Attachments
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}