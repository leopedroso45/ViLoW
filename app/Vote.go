package main

import (
	"database/sql"
	"log"
	"strconv"
)

//Vote a
type Vote struct {
	UserID, VideoID, RatingAction int
}

/*CreateVote Receives an object of type User,
opens a connection to database and returns true
if no errors occur.*/
func CreateVote(voteRep int, video Video, user User) (result bool) {
	var con *sql.DB
	con = CreateCon()

	videoID := strconv.Itoa(video.IDVideo)
	ratingAction := strconv.Itoa(voteRep)

	resultado, err := con.Query(`INSERT INTO vote (user_id, video_id, rating_action) VALUES ('` + user.IDUser + `', '` + videoID + `', '` + ratingAction + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}
