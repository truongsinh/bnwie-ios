package twitter

import (
	"smarpshare/envkey"

	"github.com/mrjones/oauth"
)

var Consumer *oauth.Consumer
var (
	mediaUploadEndpoint string
	tweetEndpoint       string
	// Have to declare it as a var for testing purpose. httptest doesn't support body > 3KB
	MaxImageSize int64
	twitterKey,
	twitterSecret string
)

func init() {
	twitterKey = envkey.String(envkey.TwitterKey)
	twitterSecret = envkey.String(envkey.TwitterSecret)
	twitterInit()
}

var twitterInit = func() {
	Consumer = oauth.NewConsumer(
		twitterKey,
		twitterSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	mediaUploadEndpoint = "https://upload.twitter.com/1.1/media/upload.json"
	tweetEndpoint = "https://api.twitter.com/1.1/statuses/update.json"
	// the max size for twitter image is 3MB. Ref: https://dev.twitter.com/rest/public/uploading-media
	MaxImageSize = 1024 * 1024 * 3
}
