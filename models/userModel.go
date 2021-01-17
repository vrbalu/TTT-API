package models

type User struct {
	Id string `json:"id,omitempty"`

	Username string `json:"username"`

	Email string `json:"emails"`

	Password string `json:"password,omitempty"`

	ExtID string `json:"extId,omitempty"`

	IDToken string `json:"idToken,omitempty"`

	Online bool `json:"online,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`
}
