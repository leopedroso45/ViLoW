package main

import (
	"database/sql"
	"fmt"
	"log"
)

/*Video object */
type Video struct {
	IDVideo                         int
	NameVideo, PathVideo, DescVideo string
}

func insertVideoIntoDB(videoA Video) (result bool) {
	var con *sql.DB
	con = CreateCon()

	resultado, err := con.Query(`INSERT INTO video (name_video, path_video, desc_video) VALUES ('` + videoA.NameVideo + `', '` + videoA.PathVideo + `', '` + videoA.DescVideo + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

func getVideoFromDB() (videoSlice []Video) {

	fmt.Println("Trying to recover videos ...")
	var con *sql.DB
	con = CreateCon()

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

func clearDB() (result bool) {
	var con *sql.DB
	con = CreateCon()

	resultado, err := con.Query("TRUNCATE TABLE video;")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}
