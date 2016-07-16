package xing

import (
	"smarpshare/envkey"

	"github.com/mrjones/oauth"
)

var (
	xingId     string
	xingSecret string
)

func init() {
	xingId = envkey.String(envkey.XingId)
	xingSecret = envkey.String(envkey.XingSecret)
	consumerInit()
}

var Consumer *oauth.Consumer

func consumerInit() {
	Consumer = oauth.NewConsumer(
		xingId,
		xingSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.xing.com/v1/request_token",
			AuthorizeTokenUrl: "https://api.xing.com/v1/authorize",
			AccessTokenUrl:    "https://api.xing.com/v1/access_token",
		})
}
