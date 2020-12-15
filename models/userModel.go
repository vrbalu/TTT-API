package models

type User struct {
	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	Online bool `json:"online,omitempty"`
}
