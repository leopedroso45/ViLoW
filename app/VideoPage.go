package main

/*VideoPage object */
type VideoPage struct {
	video Video
	user  User
}

//VideoPageConstructor asdsa
func VideoPageConstructor(video Video, user User) (videoPage VideoPage) {

	videoPage.video = video
	videoPage.user = user
	return
}
