package models

type Registration struct {
	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}
