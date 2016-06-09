package main

import "time"

func main() {

}

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

type Team struct {
	Id      int
	Country string
	Name    string // Spitzname oder so
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string

	Predictions []Prediction // Anzahl an getippten Spielen, bzw die Tipps
	Points      int
}

type Prediction struct {
	UserId     int
	MatchId    int
	GoalsHost  int  // Tore für Heimteam
	GoalsGuest int  // Tore für Gast
	Overtime   bool // evtl. wetten auf Sieg in Verlängerung, blabla
	Shootout   bool // siehe Overtime
	Result     int  // Endergebnis
}
