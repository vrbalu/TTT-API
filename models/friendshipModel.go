package models

type Friendship struct {
	Id string `json:"id,omitempty"`

	User1 string `json:"user1,omitempty"`

	User2 string `json:"user2,omitempty"`

	IsPending bool `json:"isPending,omitempty"`
}
