package main

import (
	"net/http"
	"chitchat/data"
)

func authenticate(w http.ResponseWriter, r http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			name: "_cookie",
			Value: seesion.Uuid,
			HttpOnly:true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} esle {
		http.Redirect(w, r, "/login", 302)
	}
}
