package model


import "time"

type UsersEventsAdventures struct {
	Id int64
	UserId int64 `sql:"notnull" pg:"on_delete:RESTRICT"`
	User *User
	MaximumAllowedVotes int
	EventsAdventuresId int64
	EventsAdventures *EventsAdventures `sql:"notnull" pg:"on_delete:RESTRICT"`
	Cost float64
	ExchangeId int64
	Exchange *Exchange `sql:"notnull" pg:"on_delete:RESTRICT"`
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}