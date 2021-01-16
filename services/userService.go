package services

import (
	"TTT/mod/helpers"
	"TTT/mod/models"
	"database/sql"
	"fmt"
	"log"
)

type UserServiceType struct{}

var UserService UserServiceType

const table = "dev.Users"

func init() {
	UserService = UserServiceType{}
}
func (*UserServiceType) AuthorizeUser(authorization *models.Auth) bool {
	log.Print(authorization.Email)
	query := fmt.Sprintf("SELECT Password FROM %s WHERE Email = '%s'", table, authorization.Email)
	res, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		log.Print(err)
		return false
	}
	var hash string
	if res.Next() {
		err = res.Scan(&hash)
		if err != nil {
			log.Print(err)
			return false
		}
		match := helpers.CompareHashAndPassword(hash, authorization.Password)
		if !match {
			return false
		}
		return true
	}
	return false

}
func (*UserServiceType) CheckUserExists(email string) (id string, exists bool) {
	query := fmt.Sprintf("SELECT 'ID' FROM %s WHERE Email = '%s'", table, email)
	res, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return "", false
	}
	if err != nil {
		log.Print(err)
		return "", false
	}
	res.Next()
	err = res.Scan(&id)
	if err != nil {
		log.Print(err)
		return "", false
	}
	return id, true
}
func (*UserServiceType) RegisterUserViaGoogle(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (ExtID,Username,Email,IDToken,Online,CreatedAt) VALUES ('%s','%s','%s','%s',%v,DEFAULT) ", table, user.ExtID, user.Username, user.Email, user.IDToken, user.Online)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceType) RegisterUserViaWeb(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (Username,Email,Password,Online,CreatedAt) VALUES ('%s','%s','%s',%v,DEFAULT) ", table, user.Username, user.Email, user.Password, user.Online)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceType) GetUserByEmail(email string) *models.User {
	var user models.User
	query := fmt.Sprintf("SELECT Username,Email,Online FROM %s WHERE Email = '%s'", table, email)
	res, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Print(err)
		return nil
	}
	res.Next()
	err = res.Scan(&user.Username, &user.Email, &user.Online)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &user
}
