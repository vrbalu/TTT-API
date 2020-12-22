package models

type UserData struct {
	Id        int64  `form:"id" json:"id" binding:"required"`
	Name      string `form:"name" json:"name,omitempty"`
	Email     string `form:"email" json:"email" binding:"required"`
	PhotoUrl  string `json:"photoUrl,omitempty"`
	AuthToken string `json:"authToken,omitempty"`
	IdToken   string `json:"idToken,omitempty"`
}
