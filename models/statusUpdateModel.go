package models

type StatusUpdate struct {
	Id         int  `json:"id,omitempty"`
	IsPending  bool `json:"isPending,omitempty"`
	IsFinished bool `json:"isFinished,omitempty"`
}
