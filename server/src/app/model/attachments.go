package model



type Attachments struct {
	Model
	Name string `sql:",notnull"`
	Description string
	TypeCode int `sql:",notnull"`
	Path string
	Deleatables
}