package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Running app...")

	if !clearDB() {
		log.Printf("Something went wrong cleaning db...")
	}

	if !feedDBwVideo() {
		log.Printf("Something went wrong feeding db...")
	}

	var dir string
	flag.StringVar(&dir, "dir", ".", "./web")
	flag.Parse()

	router := mux.NewRouter()
	router.PathPrefix("/data/").Handler(http.StripPrefix("/data", http.FileServer(http.Dir("./data"))))
	router.HandleFunc("/videos", GetVideo).Methods("GET")
	router.HandleFunc("/videos", PostVideo).Methods("POST")

	//router.HandleFunc("/contato/{id}", GetVideo).Methods("GET")
	//router.HandleFunc("/contato/{id}", CreateVideo).Methods("POST")
	//router.HandleFunc("/contato/{id}", DeleteVideo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

/*GetVideo to get all videos from DB*/
func GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	v := getVideoFromDB()
	json.NewEncoder(w).Encode(v)
}

/*PostVideo post a video inside a DB*/
func PostVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	r.ParseMultipartForm(1024)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("data", handler.Filename+"*.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	staticPath := "http://192.168.43.216:8000"
	currentPath := strings.Replace(tempFile.Name(), "data/", staticPath+"/data/", -1)

	videoA := VideoConstructor(handler.Filename, handler.Filename, currentPath)

	if !insertVideoIntoDB(videoA) {
		log.Println("Something went wrong adding a new video...")
		return
	}

	//fmt.Fprintf(w, "Successfully Uploaded File\n")
	w.WriteHeader(204)

}
