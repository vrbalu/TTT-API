package models

type Message struct {
	Username string `json:"username,omitempty"`

	Message string `json:"message,omitempty"`
}
