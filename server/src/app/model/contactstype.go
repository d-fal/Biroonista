package model

type Contactstype struct {
	Model
	name string
	IsPhysical bool
	IsVerifiable bool
	Typecode int /* The relation with the code */
	ParentRowId int `pg:"on_delete:RESTRICT"`
	ParentRow *Contactstype
	Deleatables
}

