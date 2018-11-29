package model

type ProvidersContacts struct {
	Model
	ContactsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Contacts *Contacts
	ProvidersId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Providers *Providers
	
	Deleatables
}