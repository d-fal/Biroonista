package model



type InvitationAux struct {
	Model
	PhysicalPath string
	AttachmentsId int64 `pg:"on_delete:RESTRICT"`
	Attachments *Attachments
	Deleatables
}