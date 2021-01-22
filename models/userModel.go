package models

type User struct {
	Id string `json:"id,omitempty"`

	Username string `json:"username"`

	Email string `json:"email"`

	Password string `json:"password,omitempty"`

	ExtID string `json:"extId,omitempty"`

	InGame bool `json:"inGame,omitempty"`

	Online bool `json:"online,omitempty"`

	RegisteredViaGoogle bool `json:"registeredViaGoogle,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`
}
