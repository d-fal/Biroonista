package model

import "time"

type EventsInvitations struct {
	Id int64
	EventsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Events *Events
	InvitationsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Invitations *Invitations
	DebtCleared bool
	Note string
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}