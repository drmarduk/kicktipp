package main

import (
	"log"
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
	_, err := OpenMatch(match)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

func OpenMatch(id int) (*Match, error) {
	stmt, err := db.Prepare("Select id, groupe, name, location, kickoff, host, guest, overtime, shootout from match where id = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	m := &Match{}
	err = row.Scan(&m.Id, &m.Group, &m.Name, &m.Location, &m.Kickoff, &m.Host, &m.Guest, &m.Overtime, &m.Shootout)
	if err != nil {
		return nil, err
	}
	return m, nil
}
