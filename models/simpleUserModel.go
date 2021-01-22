package models

type SimpleUser struct {
	Username            string `json:"username"`
	Email               string `json:"email,omitempty"`
	InGame              string `json:"inGame,omitempty"`
	Online              bool   `json:"online,omitempty"`
	RegisteredViaGoogle bool   `json:"registeredViaGoogle,omitempty"`
}
