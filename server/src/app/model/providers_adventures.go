package model



type ProvidersAdventures struct {
	
	ProvidersId int64 `sql:",pk"`
	Providers *Providers
	AdventuresId int64 `sql:",pk"`
	Adventures *Adventures
	Contacts []Contacts `pg:"many2many:user_to_contacts"`
	Deleatables
}