package models

type PlayedMove struct {
	Username string `json:"username,omitempty"`

	X string `json:"x"`

	Y string `json:"y"`

	FKGameID string `json:"game_id"`
}
