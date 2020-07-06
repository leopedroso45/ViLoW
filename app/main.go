package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("web/*.html"))

	clearDBVideo()
	clearDBUser()
	feedDBwVideo()

	log.Println("Running app...")

	router := mux.NewRouter()

	router.HandleFunc("/", IndexPageHandler).Methods("GET")
	router.HandleFunc("/sign", SignHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/login", LoginPageHandler).Methods("GET")
	router.HandleFunc("/upload", UploadPageHandler).Methods("GET")
	router.HandleFunc("/logout", LogoutHandler).Methods("POST")
	router.HandleFunc("/internal", InternalPageHandler)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./web"))))
	router.PathPrefix("/data/").Handler(http.StripPrefix("/data", http.FileServer(http.Dir("./data"))))
	router.HandleFunc("/videos", GetVideo).Methods("GET")
	router.HandleFunc("/videos", PostVideo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
