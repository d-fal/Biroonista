package model

type ProvidersServicesAttachments struct {
	Model
	ProvidersServicesId int `sql:",notnull" pg:"on_delete:RESTRICT"`
	ProvidersServices *ProvidersServices
	AttachmentsId int `sql:",notnull" pg:"on_delete:RESTRICT"`
	Attachments *Attachments
	FilePath string
	Deleatables

}