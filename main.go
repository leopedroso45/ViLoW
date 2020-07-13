package main

import (
	"ViLoW/handler"
	"ViLoW/model"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	handler.Templates = template.Must(template.ParseGlob("./web/*.html"))

	//clearDBVideo()
	//clearDBUser()
	//feedDBwVideo()

	log.Println("Running app...")

	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexPageHandler).Methods("GET")
	router.HandleFunc("/sign", handler.SignHandler).Methods("POST")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/login", handler.LoginPageHandler).Methods("GET")
	router.HandleFunc("/loging", handler.HandleGoogleLogin)
	router.HandleFunc("/callback", handler.HandleGoogleCallback)
	router.HandleFunc("/upload", handler.UploadPageHandler).Methods("POST")
	router.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")
	router.HandleFunc("/videos", model.PostVideo).Methods("POST")
	router.HandleFunc("/internal", handler.InternalPageHandler)

	router.HandleFunc("/upvote/{videoID}", handler.UpVoteHandler).Methods("POST")
	router.HandleFunc("/downvote/{videoID}", handler.DownVoteHandler).Methods("POST")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./web"))))
	router.PathPrefix("/data/").Handler(http.StripPrefix("/data", http.FileServer(http.Dir("./app/data"))))
	router.HandleFunc("/videos", model.GetVideo).Methods("GET")

	router.HandleFunc("/{id}", handler.WatchPageHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
