package models

type GameUpdate struct {
	IsPending bool `json:"isPending,omitempty"`

	IsFinished bool `json:"isFinished,omitempty"`

	GameplayX []Coordinates `json:"gameplay_x,omitempty"`

	GameplayO []Coordinates `json:"gameplay_o,omitempty"`
}
