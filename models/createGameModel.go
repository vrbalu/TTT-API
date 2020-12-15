package models

type CreateGame struct {
	Username string `json:"username,omitempty"`

	Shape string `json:"shape,omitempty"`
}
