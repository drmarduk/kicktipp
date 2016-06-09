package main

import "log"

type Prediction struct {
	UserId     int
	MatchId    int
	GoalsHost  int  // Tore für Heimteam
	GoalsGuest int  // Tore für Gast
	Overtime   bool // evtl. wetten auf Sieg in Verlängerung, blabla
	Shootout   bool // siehe Overtime
	Result     int  // Endergebnis
}

func NewPrediction(token string, user, match, goalshost, goalsguest int, overtime, shootout bool) *Prediction {
	if chk, err := checkUser(user, token); !chk && err != nil {
		log.Println("NewPrediction: User could not be found in db. " + err.Error())
		return nil
	}

	if chk, err := checkMatch(match); !chk && err != nil {
		log.Println("NewPrediction: Match could not be found in db. " + err.Error())
		return nil
	}
	return &Prediction{
		UserId:     user,
		MatchId:    match,
		GoalsHost:  goalshost,
		GoalsGuest: goalsguest,
		Overtime:   overtime,
		Shootout:   shootout,
	}
}
