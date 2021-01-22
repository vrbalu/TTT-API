package models

type GameStatsModel struct {
	WinCount int `json:"winCount,omitempty"`

	User string `json:"user,omitempty"`
}
