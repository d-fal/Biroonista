package model

type ProvidersServices struct {
	Model
	ProvidersId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Providers *Providers
	ServicesId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Services *Services
	Attachments []Attachments `pg:"many2many:providers_services_to_attachments"`
	Deleatables
}