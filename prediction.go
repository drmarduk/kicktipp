package main

import (
	"errors"
	"log"
)

type Prediction struct {
	UserId     int
	MatchId    int
	GoalsHost  int  // Tore für Heimteam
	GoalsGuest int  // Tore für Gast
	Overtime   bool // evtl. wetten auf Sieg in Verlängerung, blabla
	Shootout   bool // siehe Overtime
	Result     int  // Endergebnis
}

func NewPrediction(user, match, goalshost, goalsguest int, overtime, shootout bool) (*Prediction, error) {
	//if chk, err := validateUser(user); !chk && err != nil {
	//	log.Println("NewPrediction: User could not be found in db. " + err.Error())
	//	return nil, errors.New("user not found in db.")
	//}

	if chk, err := checkMatch(match); !chk && err != nil {
		log.Println("NewPrediction: Match could not be found in db. " + err.Error())
		return nil, errors.New("match not found in db.")
	}
	if goalshost < 0 || goalsguest < 0 {
		log.Println("NewPrediction: Goals can not be negative")
		return nil, errors.New("goalcount can not be negative")
	}
	if !overtime && shootout {
		return nil, errors.New("combination missmatch, !overtime && shootout")
	}

	return &Prediction{
		UserId:     user,
		MatchId:    match,
		GoalsHost:  goalshost,
		GoalsGuest: goalsguest,
		Overtime:   overtime,
		Shootout:   shootout,
	}, nil
}

func OpenPrediction(user, match int) (*Prediction, error) {
	return nil, nil
}

func (p *Prediction) Save() error {

	return nil
}
