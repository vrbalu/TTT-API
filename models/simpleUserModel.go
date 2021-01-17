package models

type SimpleUser struct {
	Username string `json:"username"`

	Email   string `json:"email,omitempty"`
	IdToken string `json:"idToken,omitempty"`
	Online  bool   `json:"online,omitempty"`
}
