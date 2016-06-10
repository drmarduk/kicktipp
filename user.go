package main

import (
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string

	Predictions []Prediction // Anzahl an getippten Spielen, bzw die Tipps
	Points      int
}

func checkExistingUser(name string) (bool, error) {
	stmt, err := db.Prepare("SELECT count(*) FROM user WHERE name = ?;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(name)
	var count int = 0
	if err := result.Scan(&count); err != nil {
		return false, err
	}
	if count > 1 {
		return false, nil
	}
	return true, nil
}

func OpenUser(name string) (*User, error) {
	// get user data
	stmt, err := db.Prepare("Select id, name, email, password, points from user where name = ?;")
	if err != nil {
		return nil, err
	}

	result := stmt.QueryRow(name)
	u := &User{}

	err = result.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Points)
	if err != nil {
		return nil, err
	}

	stmt.Close()
	// get predictions
	stmt, err = db.Prepare("Select userid, matchid, goalshost, goalsguest, overtime, shootout, points from prediction where userid = ?;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(u.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for !rows.Next() {
		p := Prediction{}
		err = rows.Scan(&p.UserId, &p.MatchId, &p.GoalsHost, &p.GoalsGuest, &p.Overtime, &p.Shootout, &p.Result)
		if err != nil {
			log.Println("OpenUser: Error while scanning user predictions. " + err.Error())
			continue
		}
		u.Predictions = append(u.Predictions, p)
	}

	return u, nil
}
func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("empty username")
	}
	if email == "" {
		return nil, errors.New("empty email")
	}
	if password == "" {
		return nil, errors.New("empty password")
	}
	if chk, err := checkExistingUser(name); !chk {
		if err != nil { // error while db connection foo
			return nil, err
		}
		return nil, errors.New("user already exists")
	}

	var err error
	password, err = HashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{Name: name, Email: email, Password: password}, nil
}

func (u *User) Save() error {
	stmt, err := db.Prepare("Insert into user(name, email, password) values(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}
