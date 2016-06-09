package main

type User struct {
	Id       int
	Name     string
	Email    string
	Password string

	Predictions []Prediction // Anzahl an getippten Spielen, bzw die Tipps
	Points      int
}

func checkUser(userid int, token string) (bool, error) {
	// TODO: check if userid matches token in db
	panic("not implemented")
}

func NewUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}
