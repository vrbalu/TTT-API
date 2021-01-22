package models

type UpdatePassword struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}
