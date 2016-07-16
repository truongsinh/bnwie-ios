package facebook

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var Consumer *oauth2.Config
var AppClient *http.Client
var (
	facebookId,
	facebookSecret string
)

func myinit() {
	facebookId = "envkey.String(envkey.FacebookId)"
	facebookSecret = "envkey.String(envkey.FacebookSecret)"
	consumerInit()
}

func consumerInit() {
	Consumer = &oauth2.Config{
		ClientID:     facebookId,
		ClientSecret: facebookSecret,
		Scopes:       []string{ScopePublishActions, ScopeUserFriends, ScopeUserPosts},
		RedirectURL:  "",
		Endpoint:     facebook.Endpoint,
	}
	AppClient = Consumer.Client(oauth2.NoContext, &oauth2.Token{
		AccessToken: facebookId + "|" + facebookSecret,
	})
}
