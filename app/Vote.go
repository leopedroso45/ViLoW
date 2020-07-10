package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

//Vote a
type Vote struct {
	UserID, VideoID, RatingAction int
}

/*CheckVote Receives an object of type User,
opens a connection to database and returns true
if no errors occur.*/
func CheckVote(video Video, id string) (result bool) {
	var con *sql.DB
	con = CreateCon()

	sqlst := fmt.Sprintf("SELECT COUNT(vote.rating_action) FROM vote WHERE video_id = %v and user_id = %v", video.IDVideo, id)
	row, err := con.Query(sqlst)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
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

//

/*GetLikesFromID receive the video ID */
func GetLikesFromID(id int) (likes int) {
	var con *sql.DB
	con = CreateCon()

	sqlst := fmt.Sprintf("SELECT COUNT(vote.rating_action) FROM vote WHERE rating_action = 2 and video_id = %v", id)
	row, err := con.Query(sqlst)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}

/*GetDislikesFromID receive the video ID */
func GetDislikesFromID(id int) (likes int) {
	var con *sql.DB
	con = CreateCon()

	sqlst := fmt.Sprintf("SELECT COUNT(vote.rating_action) FROM vote WHERE rating_action = 1 and video_id = %v", id)
	row, err := con.Query(sqlst)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
