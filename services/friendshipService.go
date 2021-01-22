package services

import (
	"TTT/mod/models"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type FriendshipServiceType struct{}

var FriendshipService FriendshipServiceType

const tableFriendships = "dev.Friendships"

func init() {
	FriendshipService = FriendshipServiceType{}
}

func (*FriendshipServiceType) CreateFriendship(friendship *models.FriendshipCreate) error {
	query := fmt.Sprintf("INSERT INTO %s (User1,User2,IsPending,CreatedAt) VALUES ('%s','%s',%v,DEFAULT) ", tableFriendships, friendship.User1, friendship.User2, true)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (*FriendshipServiceType) DeleteFriendship(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID = %v ", tableFriendships, id)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (*FriendshipServiceType) UpdateFriendshipPendingStatus(id string, isPending bool) error {
	query := fmt.Sprintf("UPDATE %s SET IsPending = %v WHERE ID = %s", tableFriendships, isPending, id)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func (*FriendshipServiceType) GetFriendshipById(id string) (*models.Friendship, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE ID = %s", tableFriendships, id)
	rows, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		log.Print(err)
		return nil, err
	}
	rows.Next()
	var friendship models.Friendship
	err = rows.Scan(&friendship.Id, &friendship.User1, &friendship.User2, &friendship.IsPending)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship, nil
}

func (*FriendshipServiceType) GetFriendships(user string, isPending string, forRequest string) ([]models.Friendship, error) {
	var query string
	var friendsList []models.Friendship
	boolIsPending, err := strconv.ParseBool(isPending)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	if len(user) != 0 && len(isPending) != 0 {
		query = fmt.Sprintf("SELECT ID,User1,User2,IsPending FROM %s WHERE isPending = %v AND (User1 = '%s' OR User2 = '%s') ", tableFriendships, boolIsPending, user, user)

	}
	if len(forRequest) != 0 && len(user) != 0 && len(isPending) != 0 {
		query = fmt.Sprintf("SELECT ID,User1,User2,IsPending FROM %s WHERE isPending = %v AND User2 = '%s' ", tableFriendships, boolIsPending, user)
	}
	log.Print(query)
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
		var friendship models.Friendship
		err = rows.Scan(&friendship.Id, &friendship.User1, &friendship.User2, &friendship.IsPending)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		friendsList = append(friendsList, friendship)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return friendsList, nil
}
