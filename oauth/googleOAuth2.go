package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	/*GoogleOauthConfig as*/
	GoogleOauthConfig *oauth2.Config

	// TODO: randomize it

	/*OauthStateString as*/
	OauthStateString = "pseudo-random"
)

func init() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     "123149391338-nnv2p8jvq8v7ed8jkqldp9gqr2t5lic3.apps.googleusercontent.com",
		ClientSecret: "9ECnkJbHwrHdMuEh0lJkASlQ",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
