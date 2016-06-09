package main

import "time"

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
	panic("not implemented")
}
