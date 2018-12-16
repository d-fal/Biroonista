package model



type EventsInvitations struct {
	Model
	EventsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Events *Events
	InvitationsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Invitations *Invitations
	DebtCleared bool
	Note string
	Deleatables
}