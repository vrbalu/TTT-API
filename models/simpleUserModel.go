package models

type SimpleUser struct {
	Username string `json:"username"`

	Email string `json:"email,omitempty"`

	Online bool `json:"online,omitempty"`
}
