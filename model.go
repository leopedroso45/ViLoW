package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/leopedroso45/gorest-api/db"
)

type Video struct {
	IDVideo                         int
	NameVideo, PathVideo, DescVideo string
}

func insertVideoIntoDB(videoA Video) (result bool) {
	fmt.Println("Iniciando aplicação...")
	var con *sql.DB
	con = db.CreateCon()
	nome := videoA.NameVideo
	path := videoA.PathVideo
	desc := videoA.DescVideo

	resultado, err := con.Query(`INSERT INTO video (name_video, path_video, desc_video) VALUES ('` + nome + `', '` + path + `', '` + desc + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		defer resultado.Close()
		return true
	}
}

func getVideoFromDB() (videoSlice []Video) {

	fmt.Println("Iniciando aplicação...")
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
			fmt.Println("Slice:")
			log.Println(videoSlice)
		}
	}
	return
}
