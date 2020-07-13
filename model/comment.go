package model

import (
	"ViLoW/db"
	"database/sql"
	"fmt"
	"log"
)

/*Comment object */
type Comment struct {
	IDComment, UserComment, VideoComment int
	BodyComment, DateComment             string
}

/*CommentConstructor build a video type */
func CommentConstructor(body, date string, user, video int) (comment Comment) {
	comment.BodyComment = body
	comment.DateComment = date
	comment.UserComment = user
	comment.VideoComment = video
	return comment
}

/*insertCommentIntoDB Receives an object of type Video,
opens a connection to database and returns true
if no errors occur.*/
func insertCommentIntoDB(comment Comment) (result bool) {
	var con *sql.DB
	con = db.CreateCon()
	user := string(comment.UserComment)
	video := string(comment.VideoComment)

	resultado, err := con.Query(
		`INSERT INTO comment (body_comment, user_comment, video_comment, time_comment) VALUES ('` + comment.BodyComment + `', '` + user + `', '` + video + `', '` + comment.DateComment + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

/*getCommentFromDB Receives a video index, open a connection
to database and returns a slice
 of Comments if no errors occur.*/
func getCommentFromDB(videoindex int) (commentSlice []Comment) {

	fmt.Println("Trying to recover comments ...")
	var con *sql.DB
	con = db.CreateCon()

	indexvideo := string(videoindex)

	resultado, err := con.Query("select id_comment, body_comment, user_comment, video_comment, time_comment from comment where video_comment = " + indexvideo)
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()
	for resultado.Next() {
		var comment Comment
		err := resultado.Scan(&comment.IDComment, &comment.BodyComment, &comment.UserComment, &comment.VideoComment, &comment.DateComment)
		if err != nil {
			log.Fatal(err)
		} else {
			commentSlice = append(commentSlice, comment)
		}
	}
	return
}

/*clearDBComment Open a connection
to database and clears all
comments already inserted.*/
func clearDBComment() (result bool) {
	var con *sql.DB
	con = db.CreateCon()
	resultado, err := con.Query("DELETE FROM comment;")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

/*clearDBCommentUser Open a connection
to database and clears all user
comments already inserted.*/
func clearDBCommentUser(user, comment int) (result bool) {
	var con *sql.DB
	con = db.CreateCon()
	userIndex := string(user)
	commentIndex := string(comment)

	resultado, err := con.Query("DELETE FROM comment WHERE user_comment = " + userIndex + " AND id_comment = " + commentIndex + ";")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}
