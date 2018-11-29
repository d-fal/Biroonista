package model



type Providers struct {
	Model
	Name string
	Description string
	lat float32
	long float32
	Rate int
	Services []Services `pg:"many2many:providers_services"`
	Attachments []Attachments `pg:"many2many:providers_attachments"`
	Contacts []Contacts `pg:"many2many:providers_contacts"`
	Deleatables
}