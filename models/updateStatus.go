package models

type UpdateStatus struct {
	Username string `json:"username,omitempty"`
	Online   bool   `json:"online,omitempty"`
	InGame   bool   `json:"inGame,omitempty"`
}
