package model

import "time"
type Model struct {
	Id int64
	
}

type Deleatables struct {
	CreatedAt time.Time `sql:",notnull, default:now()"`
	UpdatedAt time.Time
	DeletedAt time.Time `pg:"soft_delete"`
}