package services

import (
	"TTT/mod/models"
	"fmt"
	"log"
)

type GameServiceType struct{}

var GameService GameServiceType

const tableGames = "dev.Games"
const tableMoves = "dev.GameMoves"

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
