package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if feedDBwVideo() {
		fmt.Printf("Success feeding db...")
	} else {
		fmt.Printf("Something went wrong feeding db...")
	}
	fmt.Println("Running app...")
	router := mux.NewRouter()
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
