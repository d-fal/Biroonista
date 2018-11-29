package model


import "time"

type UsersEventsAdventuresVoteoptions struct {
	UsersEventsAdventuresId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	UsersEventsAdventures *UsersEventsAdventures 
	VoteoptionsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Voteoptions *Voteoptions 
	Comment string
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}