package models

type CreateGame struct {
	User1 string `json:"user1,omitempty"`

	User2 string `json:"user2,omitempty"`
}
