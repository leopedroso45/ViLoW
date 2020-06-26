package main

import (
	"log"
	"net/http"
)

/*IndexPageHandler as */
func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

/*SignHandler as */
func SignHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	age := r.FormValue("age")
	redirectTarget := "/"
	if name != "" && pass != "" && age != "" {

		var user User
		user.NameUser = name
		user.PasswordUser = pass
		user.AgeUser = age
		CreateUser(user)
		user, err := GetUser(name, pass)
		if err != nil {
			log.Println(err)
			return
		}
		SetSession(user.IDUser, user.NameUser, w)
		redirectTarget = "/internal"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

/*LoginHandler as */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("user")
	pass := r.FormValue("pass")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		user, err := GetUser(name, pass)
		if err == nil {
			SetSession(user.IDUser, user.NameUser, w)
			redirectTarget = "/internal"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

/*LogoutHandler as */
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSession(w)
	http.Redirect(w, r, "/", 302)
}

/*InternalPageHandler as */
func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, _ := GetUserName(r)
	if userName != "" {
		templates.ExecuteTemplate(w, "home.html", userName)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
