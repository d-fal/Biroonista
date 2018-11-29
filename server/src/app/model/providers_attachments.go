package model



type ProvidersAttachments struct {
	Model
	ProvidersId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Providers *Providers
	AttachmentsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Attachments *Attachments
	Deleatables
}