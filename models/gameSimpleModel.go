package models

type GameSimple struct {
	Id string `json:"id,omitempty"`

	UserX string `json:"user_x,omitempty"`

	UserO string `json:"user_o,omitempty"`

	IsPending bool `json:"isPending,omitempty"`

	IsFinished bool `json:"isFinished,omitempty"`
}
