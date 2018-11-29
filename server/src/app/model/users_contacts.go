package model


type UserContact struct {
	Model
	UserId int `sql:",pk"`
	User *User
	ContactsId int `sql:",pk"`
	Contacts *Contacts
	Deleatables
}