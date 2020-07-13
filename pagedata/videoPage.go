package pagedata

import (
	"ViLoW/model"
)

/*VideoPage object */
type VideoPage struct {
	Video                        model.Video
	User                         model.User
	PositiveCount, NegativeCount int
	HasVote                      bool
}

//VideoPageConstructor asdsa
func VideoPageConstructor(video model.Video, user model.User, pos, neg int, hasVote bool) (videoPage VideoPage) {

	videoPage.Video = video
	videoPage.User = user
	videoPage.PositiveCount = pos
	videoPage.NegativeCount = neg
	videoPage.HasVote = hasVote
	return
}

//VideoPageConstructor2 asdsa
func VideoPageConstructor2(video model.Video, user model.User, pos, neg int) (videoPage VideoPage) {

	videoPage.Video = video
	videoPage.User = user
	videoPage.PositiveCount = pos
	videoPage.NegativeCount = neg
	return
}
