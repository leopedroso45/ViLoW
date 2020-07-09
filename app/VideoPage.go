package main

/*VideoPage object */
type VideoPage struct {
	Video Video
	User  User
}

//VideoPageConstructor asdsa
func VideoPageConstructor(video Video, user User) (videoPage VideoPage) {

	videoPage.Video = video
	videoPage.User = user
	return
}
