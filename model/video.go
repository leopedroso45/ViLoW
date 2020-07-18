package model

import (
	"ViLoW/db"
	"ViLoW/session"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*Video object */
type Video struct {
	IDVideo, IDUser                 int
	NameVideo, PathVideo, DescVideo string
}

/*VideoConstructor build a video type */
func VideoConstructor(name, desc, path string) (video Video) {
	video.NameVideo = name
	video.PathVideo = path
	video.DescVideo = desc
	return video
}

/*insertVideoIntoDB Receives an object of type Video,
opens a connection to database and returns true
if no errors occur.*/
func insertVideoIntoDB(videoA Video) (result bool) {
	var con *sql.DB
	con = db.CreateCon()

	resultado, err := con.Query(`INSERT INTO video (name_video, path_video, desc_video) VALUES ('` + videoA.NameVideo + `', '` + videoA.PathVideo + `', '` + videoA.DescVideo + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

/*insertVideoIntoDBWID Receives an object of type Video and the USER ID,
opens a connection to database and returns true
if no errors occur.*/
func insertVideoIntoDBWID(videoA Video, id string) (result bool) {
	var con *sql.DB
	con = db.CreateCon()

	resultado, err := con.Query(`INSERT INTO video (usuario_video, name_video, path_video, desc_video) VALUES ('` + id + `', '` + videoA.NameVideo + `', '` + videoA.PathVideo + `', '` + videoA.DescVideo + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

/*GetAVideoFromDB Open a connection
to database and returns a Video if no errors occur.*/
func GetAVideoFromDB(id string) (video Video) {
	var con *sql.DB
	con = db.CreateCon()

	resultado, err := con.Query(`SELECT id_video, usuario_video, name_video, path_video, desc_video from video where id_video = '` + id + `'`)
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()
	for resultado.Next() {
		err := resultado.Scan(&video.IDVideo, &video.IDUser, &video.NameVideo, &video.PathVideo, &video.DescVideo)
		if err != nil {

			log.Fatal(err)
			return
		}
	}
	return
}

/*GetVideo6FromDB Open a connection
to database and returns a slice
 of Video if no errors occur.*/
func GetVideo6FromDB() (videoSlice []Video) {

	fmt.Println("Trying to recover videos ...")
	var con *sql.DB
	con = db.CreateCon()

	resultado, err := con.Query("SELECT id_video,name_video, path_video, desc_video from video ORDER BY RAND() LIMIT 6")
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()

	for resultado.Next() {
		var videoA Video
		err := resultado.Scan(&videoA.IDVideo, &videoA.NameVideo, &videoA.PathVideo, &videoA.DescVideo)

		if err != nil {
			log.Fatal(err)
		} else {
			videoSlice = append(videoSlice, videoA)
		}
	}
	return
}

/*GetVideoFromDB Open a connection
to database and returns a slice
 of Video if no errors occur.*/
func GetVideoFromDB() (videoSlice []Video) {

	fmt.Println("Trying to recover videos ...")
	var con *sql.DB
	con = db.CreateCon()
	resultado, err := con.Query("select id_video,name_video, path_video, desc_video from video")
	if err != nil {

		log.Fatal(err)
	}
	defer resultado.Close()

	for resultado.Next() {
		var videoA Video
		err := resultado.Scan(&videoA.IDVideo, &videoA.NameVideo, &videoA.PathVideo, &videoA.DescVideo)
		if err != nil {

			log.Fatal(err)
		} else {
			videoSlice = append(videoSlice, videoA)
		}
	}
	return
}

/*ClearDBVideo Open a connection
to database and clears all
videos already inserted.*/
func ClearDBVideo() (result bool) {
	var con *sql.DB
	con = db.CreateCon()

	resultado, err := con.Query("DELETE FROM video;")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

/*GetVideo to get all videos from DB*/
func GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	v := GetVideoFromDB()
	json.NewEncoder(w).Encode(v)
}

/*PostVideo post a video inside a DB*/
func PostVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	name := r.FormValue("name")
	desc := r.FormValue("desc")

	r.ParseMultipartForm(1024)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("./app/data", handler.Filename+"*.mp4")
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

	//staticPath, _ := getLocalIP()
	//port := ":8000"
	//staticPath = staticPath + port

	staticPath := "http://192.168.43.216:8000"
	currentPath := strings.Replace(tempFile.Name(), "app/data/", staticPath+"/data/", -1)

	videoA := VideoConstructor(name, desc, currentPath)

	_, id := session.GetUserName(r)
	if !insertVideoIntoDBWID(videoA, id) {
		log.Println("Something went wrong adding a new video...")
		return
	}

	//fmt.Fprintf(w, "Successfully Uploaded File\n")
	//templates.ExecuteTemplate(w, "home.html", userName)
	http.Redirect(w, r, "/internal", 302)

}
