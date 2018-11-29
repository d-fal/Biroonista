package model

import (
	"time"
)
type Events struct {
	Id int64
	Subject string
	Description string
	StartAt time.Time
	EndAt time.Time
	IsPublic bool
	MaxInvitees int
	MinInvitees int
	Deadline time.Time	
	UserID int64 `sql:"on_delete:RESTRICT, notnull"`
	User *User
	IsTreated bool
	TreatShare float32
	IsOpenBudget bool
	Budget float64
	ExchangeId int64
	Exchange *Exchange
	Invitations []Invitations `pg:"many2many:events_to_invitations"`
	Adventures []Adventures `pg:"many2many:events_to_adventures"`
	
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}

