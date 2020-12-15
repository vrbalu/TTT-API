package models

type PlayedMove struct {
	Username string `json:"username,omitempty"`

	Shape string `json:"shape,omitempty"`

	Coor []Coordinates `json:"coor,omitempty"`
}
