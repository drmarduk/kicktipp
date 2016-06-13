package main

import (
	"net/http"
	"time"
)

var currentUser *User

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("login")
		if err == http.ErrNoCookie || c.Value == "" {
			// not logged in
			login(w, r)
			return
		}
		if c.Expires.Before(time.Now()) {
			login(w, r)
			return
		}

		// check cookie in db
		currentUser, err = OpenUserBySession(c.Value)
		if err != nil {
			login(w, r)
			return
		}

		// nice, go to page we want
		h.ServeHTTP(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
}

func errorHandler(w http.ResponseWriter, r *http.Request, msg string) {
	w.Write([]byte("Error while working: " + msg))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the members area " + currentUser.Name))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello " + currentUser.Name))
	return
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		// process login
		user := r.FormValue("user")
		pass := r.FormValue("pass")

		currentUser, err = OpenUserByName(user)
		if err != nil {
			errorHandler(w, r, "LoginHandler(1): "+err.Error())
			return
		}

		chk, err := ComparePassword(pass, currentUser.Password)
		if err != nil {
			errorHandler(w, r, "LoginHandler(2): "+err.Error())
			return
		}
		if !chk {
			errorHandler(w, r, "LoginHandler(3): "+"wrong pw")
			return
		}

		if err = currentUser.GenerateSession(); err != nil {
			errorHandler(w, r, "LoginHandler(4): "+err.Error())
			return
		}
		setCookie(w, currentUser.Session)
		w.Write([]byte("welcome " + currentUser.Name))
		return

	} else {
		// print login page
		w.Write([]byte(htmllogin))
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		user := r.FormValue("user")
		pass := r.FormValue("pass")
		email := r.FormValue("email")

		if user == "" || pass == "" {
			errorHandler(w, r, "user and pass must not be empty")
			return
		}
		currentUser, err = NewUser(user, email, pass)
		if err != nil {
			errorHandler(w, r, "RegisterHandler(1): "+err.Error())
			return
		}
		err = currentUser.Save()
		if err != nil {
			errorHandler(w, r, "RegisterHandler(2): "+err.Error())
			return
		}
		if err = currentUser.GenerateSession(); err != nil {
			errorHandler(w, r, "RegisterHandler(3): "+err.Error())
			return
		}

		setCookie(w, currentUser.Session)
		http.Redirect(w, r, "/main", 302)
	} else {
		w.Write([]byte(htmlregister))
	}

}

func setCookie(w http.ResponseWriter, session string) {
	c := &http.Cookie{Name: "session", Value: session, Expires: time.Now().Add(time.Hour * 2)}
	http.SetCookie(w, c)
	return
}
