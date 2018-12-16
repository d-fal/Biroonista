package model


type EventsUsers struct {
	Model
	EventsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Events *Events 
	UsersId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Users *Users
	Remark string
	Deleatables
}