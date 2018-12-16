package model

import "time"
type Events struct {
	Model
	Subject string
	Description string
	StartAt time.Time
	EndAt time.Time
	IsPublic bool
	MaxInvitees int
	MinInvitees int
	Deadline time.Time	
	UsersID int64 `sql:"on_delete:RESTRICT, notnull"`
	Users *Users
	IsTreated bool
	TreatShare float32
	IsOpenBudget bool
	Budget float64
	ExchangeId int64
	Exchange *Exchange
	Invitations []Invitations `pg:"many2many:events_to_invitations"`
	Adventures []Adventures `pg:"many2many:events_to_adventures"`
	Deleatables
}

