package models

type Game struct {
	Id string `json:"id,omitempty"`

	UserX string `json:"user_x,omitempty"`

	UserO string `json:"user_o,omitempty"`

	IsPending bool `json:"isPending,omitempty"`

	IsFinished bool `json:"isFinished,omitempty"`

	GameplayX []Coordinates `json:"gameplay_x,omitempty"`

	GameplayO []Coordinates `json:"gameplay_o,omitempty"`
}
