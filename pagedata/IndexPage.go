package pagedata

import (
	"ViLoW/model"
)

/*IndexPage means the index object*/
type IndexPage struct {
	User   model.User
	Videos []model.Video
	Logged bool
}

/*IndexPageVideoConstructor means the index object*/
func IndexPageVideoConstructor(videos []model.Video) (page IndexPage) {
	page.Videos = videos
	page.Logged = false
	return
}

/*IndexPageConstructorU means the index object*/
func IndexPageConstructorU(user model.User, videos []model.Video) (page IndexPage) {
	page.User = user
	page.Videos = videos
	page.Logged = true
	return
}
