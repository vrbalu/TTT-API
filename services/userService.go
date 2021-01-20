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
	query := fmt.Sprintf("INSERT INTO %s (ExtID,Username,Email,InGame,Online,RegisteredViaGoogle,CreatedAt) VALUES ('%s','%s','%s','%s',%v,%v,DEFAULT) ", table, user.ExtID, user.Username, user.Email, user.InGame, user.Online, user.RegisteredViaGoogle)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceType) RegisterUserViaWeb(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (Username,Email,Password,InGame,Online,RegisteredViaGoogle,CreatedAt) VALUES ('%s','%s','%s',%v,%v,%v,DEFAULT) ", table, user.Username, user.Email, user.Password, user.InGame, user.Online, user.RegisteredViaGoogle)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceType) GetUserByEmail(email string) *models.User {
	var user models.User
	query := fmt.Sprintf("SELECT Username,Email,InGame,Online,RegisteredViaGoogle FROM %s WHERE Email = '%s'", table, email)
	res, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Print(err)
		return nil
	}
	res.Next()
	err = res.Scan(&user.Username, &user.Email, &user.InGame, &user.Online, &user.RegisteredViaGoogle)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &user
}
func (*UserServiceType) UpdateStatus(email string, statusField string, updatedValue bool) error {
	query := fmt.Sprintf("UPDATE %s SET %s = %v WHERE Email = '%s'", table, statusField, updatedValue, email)
	log.Print("Updating status")
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (*UserServiceType) UpdatePassword(email string, updatedValue string) error {
	query := fmt.Sprintf("UPDATE %s SET Password = '%s' WHERE Email = '%s'", table, updatedValue, email)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (*UserServiceType) GetAllUsers(onlyOnline bool) ([]models.SimpleUser, error) {
	var activeUsers []models.SimpleUser
	var query string
	if onlyOnline {
		query = fmt.Sprintf("SELECT Username,Email,InGame,Online,RegisteredViaGoogle FROM %s WHERE Online = true", table)
	} else {
		query = fmt.Sprintf("SELECT Username,Email,InGame,Online,RegisteredViaGoogle FROM %s", table)
	}
	rows, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.SimpleUser
		err = rows.Scan(&user.Username, &user.Email, &user.InGame, &user.Online, &user.RegisteredViaGoogle)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		activeUsers = append(activeUsers, user)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return activeUsers, nil
}
