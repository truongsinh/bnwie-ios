package twitter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrjones/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwitter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Twitter Suite")
}

var _ = Describe("Twitter", func() {
	accessToken := oauth.AccessToken{
		Token:  "twitterToken",
		Secret: "twitterSecret",
	}
	Context("UploadMedia test", func() {
		It("can parse the media_id_string field in response when 200 OK", func() {
			imageServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				fmt.Fprintln(w, "image_binary_data")
			}))
			twitterServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				fmt.Fprintln(w, `{"media_id":611490910875058177,"media_id_string":"611490910875058177","size":625539,"expires_after_secs":3600,"image":{"image_type":"image\/jpeg","w":2560,"h":1440}}`)
			}))
			mediaUploadEndpoint = twitterServer.URL
			defer imageServer.Close()
			defer twitterServer.Close()
			media_id, err := UploadMedia(imageServer.URL, accessToken)
			Expect(media_id).To(Equal("611490910875058177"))
			Expect(err).To(BeNil())
		})
		It("should not call postMedia and return empty string and err when fetch image fails", func() {
			called := false
			postMedia = func(dataReader io.Reader, token *oauth.AccessToken) (media_id string, err error) {
				called = true
				return
			}
			imageServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				fmt.Fprintln(w, "what ever error message")
			}))
			defer imageServer.Close()
			media_id, err := UploadMedia(imageServer.URL, accessToken)
			Expect(called).To(Equal(false))
			Expect(media_id).To(Equal(""))
			Expect(err).NotTo(BeNil())
		})
		It("return empty string and UploadTooLargeImageError when image is too large", func() {
			MaxImageSize_orig := MaxImageSize
			// Content-Length will be the true length of the body in response, no matter what Content-Length you set in request, so we have to create the body of the exact length
			// weird, when body is larger than 1024*3, Content-Length : -1 in response. So we need to use small MaxImageSize in test
			MaxImageSize = 1024 * 2
			imageServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				slice := make([]byte, 1024*2)
				for i := range slice {
					slice[i] = byte('A')
				}
				w.Write(slice)
			}))
			defer imageServer.Close()
			media_id, err := UploadMedia(imageServer.URL, accessToken)
			Expect(media_id).To(Equal(""))
			Expect(err).To(Equal(UploadTooLargeImageError))
			MaxImageSize = MaxImageSize_orig
		})
		It("return empty string and err when twitter responds not 200 OK", func() {
			imageServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				fmt.Fprintln(w, "image_binary_data")
			}))
			twitterServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				w.Header().Set("Content-Length", "23")
				fmt.Fprintln(w, "what ever error message")
			}))
			mediaUploadEndpoint = twitterServer.URL
			defer imageServer.Close()
			defer twitterServer.Close()
			media_id, err := UploadMedia("what ever stirng", accessToken)
			Expect(media_id).To(Equal(""))
			Expect(err).NotTo(BeNil())
		})
	})
	Context("UpdateStatuses test", func() {

		It("can parse the id_str field in response when 200 OK", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				fmt.Fprintln(w, `{"created_at":"Thu Jun 18 11:09:30 +0000 2015","id":611490916524761088,"id_str":"611490916524761088","text":"Error handling and Go http:\/\/t.co\/2cgOepyeDM #smarper http:\/\/t.co\/b2s0Ej413c","source":"\u003ca href=\"http:\/\/smarpshare.com\" rel=\"nofollow\"\u003eSmarp\u003c\/a\u003e","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"user":{"id":3040355272,"id_str":"3040355272"},"geo":null,"coordinates":null,"place":null,"contributors":null,"is_quote_status":false,"retweet_count":0,"favorite_count":0,"entities":{"hashtags":[{"text":"smarper","indices":[45,53]}],"symbols":[],"user_mentions":[],"urls":[{"url":"http:\/\/t.co\/2cgOepyeDM","expanded_url":"http:\/\/local.smh.re\/4I","display_url":"local.smh.re\/4I","indices":[22,44]}],"media":[{"id":611490910875058177,"id_str":"611490910875058177","indices":[54,76],"media_url":"http:\/\/pbs.twimg.com\/media\/CHxzvl4WgAEd1Z0.jpg","media_url_https":"https:\/\/pbs.twimg.com\/media\/CHxzvl4WgAEd1Z0.jpg","url":"http:\/\/t.co\/b2s0Ej413c","display_url":"pic.twitter.com\/b2s0Ej413c","expanded_url":"http:\/\/twitter.com\/linzhiqi07\/status\/611490916524761088\/photo\/1","type":"photo","sizes":{"thumb":{"w":150,"h":150,"resize":"crop"},"small":{"w":340,"h":191,"resize":"fit"},"medium":{"w":600,"h":337,"resize":"fit"},"large":{"w":1024,"h":576,"resize":"fit"}}}]},"extended_entities":{"media":[{"id":611490910875058177,"id_str":"611490910875058177","indices":[54,76],"media_url":"http:\/\/pbs.twimg.com\/media\/CHxzvl4WgAEd1Z0.jpg","media_url_https":"https:\/\/pbs.twimg.com\/media\/CHxzvl4WgAEd1Z0.jpg","url":"http:\/\/t.co\/b2s0Ej413c","display_url":"pic.twitter.com\/b2s0Ej413c","expanded_url":"http:\/\/twitter.com\/linzhiqi07\/status\/611490916524761088\/photo\/1","type":"photo","sizes":{"thumb":{"w":150,"h":150,"resize":"crop"},"small":{"w":340,"h":191,"resize":"fit"},"medium":{"w":600,"h":337,"resize":"fit"},"large":{"w":1024,"h":576,"resize":"fit"}}}]},"favorited":false,"retweeted":false,"possibly_sensitive":false,"lang":"en"}`)
			}))
			defer ts.Close()
			tweetEndpoint = ts.URL
			tweet_id, err := UpdateStatuses("my status", "1234", accessToken)
			Expect(tweet_id).To(Equal("611490916524761088"))
			Expect(err).To(BeNil())
		})
		It("return empty string and err when not 200 OK", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(403)
				w.Header().Set("Content-Length", "68")
				fmt.Fprintln(w, `{"errors":[{"code":186,"message":"Status is over 140 characters."}]}`)
			}))
			defer ts.Close()
			tweetEndpoint = ts.URL
			tweet_id, err := UpdateStatuses("my status", "1234", accessToken)
			Expect(tweet_id).To(Equal(""))
			Expect(err).NotTo(BeNil())
			//below is specific for 403-186 error
			Expect(err.Error()).To(Equal("Status is over 140 characters"))
		})
	})
})
