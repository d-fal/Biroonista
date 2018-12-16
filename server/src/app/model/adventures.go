package model

type Adventures struct {
	Id int64
	Name string `sql:",notnull"`
	Description string
	FlowsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Deleatables
}