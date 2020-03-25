package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Running app...")

	//time.Sleep(10 * time.Second)

	if feedDBwVideo() {
	} else {
		fmt.Printf("Something went wrong feeding db...")
	}

	var dir string

	flag.StringVar(&dir, "dir", ".", "/web")
	flag.Parse()

	router := mux.NewRouter()
	// This will serve files under http://localhost:8000/static/<filename>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	router.HandleFunc("/videos", GetVideo).Methods("GET")
	/**
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	*/
	log.Fatal(http.ListenAndServe(":8000", router))
}

/*GetVideo to get all videos from DB*/
func GetVideo(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	v := getVideoFromDB()
	json.NewEncoder(w).Encode(v)

}
