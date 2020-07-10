package main

/*VideoPage object */
type VideoPage struct {
	Video                        Video
	User                         User
	PositiveCount, NegativeCount int
	HasVote                      bool
}

//VideoPageConstructor asdsa
func VideoPageConstructor(video Video, user User, pos, neg int, hasVote bool) (videoPage VideoPage) {

	videoPage.Video = video
	videoPage.User = user
	videoPage.PositiveCount = pos
	videoPage.NegativeCount = neg
	videoPage.HasVote = hasVote
	return
}

//VideoPageConstructor2 asdsa
func VideoPageConstructor2(video Video, user User, pos, neg int) (videoPage VideoPage) {

	videoPage.Video = video
	videoPage.User = user
	videoPage.PositiveCount = pos
	videoPage.NegativeCount = neg
	return
}
