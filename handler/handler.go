package handler

import (
	"ViLoW/model"
	"ViLoW/pagedata"
	"ViLoW/session"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

/*Templates var */
var Templates *template.Template

/*IndexPageHandler as */
func IndexPageHandler(w http.ResponseWriter, r *http.Request) {

	userName, id := session.GetUserName(r)
	if userName == "" && id == "" {
		videos := model.GetVideo6FromDB()
		res := pagedata.IndexPageVideoConstructor(videos)
		Templates.ExecuteTemplate(w, "index.html", res)
		return
	} else {
		//refatore
		videos := model.GetVideoFromDB()
		idu, _ := strconv.Atoi(id)
		user, err := model.GetUserFromID(idu)
		if err == nil {
			res := pagedata.IndexPageConstructorU(user, videos)
			Templates.ExecuteTemplate(w, "index.html", res)
			return
		}
		//refatore

	}
}

/*LoginPageHandler as */
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "login.html", nil)
}

/*UploadPageHandler as */
func UploadPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, _ := session.GetUserName(r)
	if userName != "" {
		Templates.ExecuteTemplate(w, "upload.html", nil)
	} else {
		http.Redirect(w, r, "/", 302)

	}

}

/*VoteHandler as*/
func VoteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["videoID"]
	videoID, _ := strconv.Atoi(id)
	updated := model.GetUpdatedData(videoID)

	json.NewEncoder(w).Encode(updated)
	return
}

/*UpVoteHandler as*/
func UpVoteHandler(w http.ResponseWriter, r *http.Request) {
	userName, id := session.GetUserName(r)
	if userName == "" && id == "" {
		Templates.ExecuteTemplate(w, "login.html", userName)

	} else {
		vars := mux.Vars(r)
		idd := vars["videoID"]
		video := model.GetAVideoFromDB(idd)
		idu, _ := strconv.Atoi(id)
		user, err := model.GetUserFromID(idu)
		if err == nil {
			model.CreateVote(2, video, user)
			http.Redirect(w, r, "/internal", 302)
			return
		}
	}
	return

}

/*DownVoteHandler as*/
func DownVoteHandler(w http.ResponseWriter, r *http.Request) {
	userName, id := session.GetUserName(r)
	if userName == "" && id == "" {
		Templates.ExecuteTemplate(w, "login.html", userName)

	} else {
		vars := mux.Vars(r)
		idd := vars["videoID"]
		video := model.GetAVideoFromDB(idd)
		idu, _ := strconv.Atoi(id)
		user, err := model.GetUserFromID(idu)
		if err == nil {
			model.CreateVote(1, video, user)
			http.Redirect(w, r, "/internal", 302)
			return
		}
	}
	return
}

/*WatchPageHandler as*/
func WatchPageHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	video := model.GetAVideoFromDB(vars["id"])
	likes := model.GetLikesFromID(video.IDVideo)
	dislikes := model.GetDislikesFromID(video.IDVideo)
	user, err := model.GetUserFromID(video.IDUser)
	if err == nil {
		userName, id := session.GetUserName(r)
		if userName != "" && id != "" {
			hasVote := model.CheckVote(video, id)
			page := pagedata.VideoPageConstructor(video, user, likes, dislikes, hasVote)
			Templates.ExecuteTemplate(w, "watch.html", page)

		} else {
			page := pagedata.VideoPageConstructor2(video, user, likes, dislikes)
			Templates.ExecuteTemplate(w, "watch.html", page)
		}

	} else {
		log.Print(err)
		Templates.ExecuteTemplate(w, "index.html", nil)
	}

}

/*SignHandler as */
func SignHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	age := r.FormValue("age")
	redirectTarget := "/"
	if name != "" && pass != "" && age != "" {

		var user model.User
		user.NameUser = name
		user.PasswordUser = pass
		user.AgeUser = age
		model.CreateUser(user)
		user, err := model.GetUser(name, pass)
		if err != nil {
			log.Println(err)
			return
		}
		session.SetSession(user.IDUser, user.NameUser, w)
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
		user, err := model.GetUser(name, pass)
		if err == nil {
			session.SetSession(user.IDUser, user.NameUser, w)
			redirectTarget = "/internal"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

/*LogoutHandler as */
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}

/*InternalPageHandler as */
func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	userName, _ := session.GetUserName(r)
	if userName != "" {
		var user model.UserG
		user.Name = userName
		Templates.ExecuteTemplate(w, "home.html", user)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
