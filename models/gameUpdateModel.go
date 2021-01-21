package models

type GameUpdate struct {
	Id int `json:"id,omitempty"`

	Winner string `json:"winner,omitempty"`

	IsPending bool `json:"isPending,omitempty"`

	IsFinished bool `json:"isFinished,omitempty"`
}
