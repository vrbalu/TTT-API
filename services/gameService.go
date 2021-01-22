package services

import (
	"TTT/mod/models"
	"database/sql"
	"fmt"
	"log"
)

type GameServiceType struct{}

var GameService GameServiceType

const tableGames = "dev.Games"

func init() {
	GameService = GameServiceType{}
}

func (*GameServiceType) CreateGame(newGame *models.CreateGame) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (User1,User2,IsPending,IsFinished,Winner,CreatedAt) VALUES ('%s','%s',true,false,'%s',DEFAULT) ", tableGames, newGame.User1, newGame.User2, "")
	log.Print(query)
	res, err := DbService.Exec(query)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), nil
}

func (*GameServiceType) UpdateGame(id int, winner string, isPending bool, isFinished bool) error {
	query := fmt.Sprintf("UPDATE %s SET Winner = '%s', IsPending = %v, IsFinished = %v WHERE id = %v", tableGames, winner, isPending, isFinished, id)
	log.Print(query)
	_, err := DbService.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (*GameServiceType) GetGameStats() ([]models.GameStatsModel, error) {
	var resultArr []models.GameStatsModel
	query := fmt.Sprintf("SELECT COUNT(ID),Winner FROM %s GROUP BY Winner ORDER BY COUNT(ID) DESC", tableGames)
	rows, err := DbService.Query(query)
	if err == sql.ErrNoRows {
		return resultArr, err
	}
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var result models.GameStatsModel
		err = rows.Scan(&result.WinCount, &result.User)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if result.User != "" {
			resultArr = append(resultArr, result)
		}
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return resultArr, nil

}
