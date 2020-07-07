package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*IndexPageHandler as */
func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	videos := getVideo6FromDB()
	templates.ExecuteTemplate(w, "index.html", videos)
}

/*LoginPageHandler as */
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

/*UploadPageHandler as */
func UploadPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, _ := GetUserName(r)
	if userName != "" {
		templates.ExecuteTemplate(w, "upload.html", nil)
	} else {
		http.Redirect(w, r, "/", 302)

	}

}

/*WatchPageHandler as */
func WatchPageHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	video := getAVideoFromDB(vars["id"])
	user, err := GetUserFromID(string(video.IDUser))
	if err == nil {
		page := VideoPageConstructor(video, user)
		templates.ExecuteTemplate(w, "watch.html", page)
	} else {
		log.Print(err)
		templates.ExecuteTemplate(w, "index.html", nil)
	}

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
	fmt.Printf("DEU MERDA")
}

/*InternalPageHandler as */
func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, _ := GetUserName(r)
	if userName != "" {
		fmt.Printf("DEU MERDA")
		templates.ExecuteTemplate(w, "home.html", userName)
	} else {
		fmt.Printf("DEU MERDA")
		http.Redirect(w, r, "/", 302)
	}
}
