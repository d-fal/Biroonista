package model


const modelname = "Exchange"

type Exchange struct{
	Model
	Name string `sql:"unique"`
	fx float64 `sql:",notnull"`
	Deleatables
}
