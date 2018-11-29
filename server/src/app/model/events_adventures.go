package model


import "time"

type EventsAdventures struct {
	Id int64
	AdventuresId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Adventures *Adventures
	StartAt time.Time
	EndAt time.Time
	Capacity int
	IsVoted bool
	ProvidersId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Providers *Providers
	User []User `pg:"many2many:events_adventures_to_users"`
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}