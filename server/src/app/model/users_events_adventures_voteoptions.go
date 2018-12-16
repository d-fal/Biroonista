package model


type UsersEventsAdventuresVoteoptions struct {
	Model
	UsersEventsAdventuresId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	UsersEventsAdventures *UsersEventsAdventures 
	VoteoptionsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Voteoptions *Voteoptions 
	Comment string
	Deleatables
}