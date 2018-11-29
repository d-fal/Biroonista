package model

type FlowProcedures struct {
	Model
	FlowsId int64 `sql:",notnull" pg:"on_delete:RESTRICT"`
	Flows *Flows
	Name string
	Priority int
	PendingStatusId int
	NextStatusIfSuccessId int
	NextStatusIfCancelledId int
	NextStatusIfTerminatedId int
	NextStepIfSuccessId int
	NextStepIfSuccess *FlowProcedures
	NextStepIfCancelledId int
	NextStepIfCancelled *FlowProcedures
	FilePath string
	Deleatables

}