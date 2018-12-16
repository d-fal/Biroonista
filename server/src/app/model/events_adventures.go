package model
import "time"

type EventsAdventures struct {
	Model
	AdventuresId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Adventures *Adventures
	StartAt time.Time
	EndAt time.Time
	Capacity int
	IsVoted bool
	ProvidersId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Providers *Providers
	Users []Users `pg:"many2many:events_adventures_to_users"`
	Deleatables	
}