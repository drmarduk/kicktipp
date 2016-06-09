package main

import (
	"errors"
	"time"
)

type Match struct {
	Id       int
	Group    string    // geht dann nicht mehr in den KO Spielen
	Name     string    // 1. Vorrunde, Halbfinale, Finale, ....
	Location string    // Austragungsort
	Kickoff  time.Time // Anpfiff
	Host     Team      // Heimmannschaft
	Guest    Team      // Gastmannschaft
	Overtime bool      // gab es eine Verlängerung
	Shootout bool      // gab es ein Elfmeterschießen
}

func checkMatch(match int) (bool, error) {
	// checks wether the match is in the db
	stmt, err := db.Prepare("select id from match where id = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(match)
	var resultid int
	err = row.Scan(&resultid)
	if err != nil {
		return false, err
	}
	if resultid != match {
		return false, errors.New("matchid missmatch")
	}
	return true, nil
}
