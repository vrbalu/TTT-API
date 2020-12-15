package models

type UpdatePassword struct {
	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`
}
