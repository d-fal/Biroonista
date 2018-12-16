package model


type UsersEventsAdventures struct {
	Model
	UsersId int64 `sql:"notnull" pg:"on_delete:RESTRICT"`
	Users *Users
	MaximumAllowedVotes int
	EventsAdventuresId int64
	EventsAdventures *EventsAdventures `sql:"notnull" pg:"on_delete:RESTRICT"`
	Cost float64
	ExchangeId int64
	Exchange *Exchange `sql:"notnull" pg:"on_delete:RESTRICT"`
	Deleatables
}