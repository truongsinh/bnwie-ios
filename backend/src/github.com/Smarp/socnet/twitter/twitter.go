package twitter

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"smarpshare/internalerror"
	"smarpshare/unsafeclient"
	"strings"

	"github.com/mrjones/oauth"
	"github.com/smarp/go/encoding/json"
)

const UploadTooLargeImageError internalerror.InternalError = "upload image too large"

var fetchMedia = func(mediaUrl string) (reader io.ReadCloser, err error) {
	mediaResponse, err := unsafeclient.Get(mediaUrl)
	if err != nil {
		return nil, fmt.Errorf("UploadMedia: Cannot read %q. Error: %s", mediaUrl, err)
	}

	if mediaResponse.StatusCode != http.StatusOK {
		defer mediaResponse.Body.Close()
		return nil, fmt.Errorf("UploadMedia: Fetching image %q failed. Status code: %d", mediaUrl, mediaResponse.StatusCode)
	}
	if mediaResponse.ContentLength >= MaxImageSize {
		defer mediaResponse.Body.Close()
		// should throw specific error here
		return nil, UploadTooLargeImageError
	}
	return mediaResponse.Body, nil
}

var postMedia = func(dataReader io.Reader, token *oauth.AccessToken) (media_id string, err error) {
	response, err := Consumer.PostMultipart(mediaUploadEndpoint, "media", dataReader, nil, token)
	if err != nil {
		return "", fmt.Errorf("UploadMedia: Post failed. Error: %s", err)
	}
	// if 200 OK returned
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	var TwitterResponse struct {
		MediaId string `json:"media_id_string"`
	}
	err = decoder.Decode(&TwitterResponse)
	if err != nil {
		return "", fmt.Errorf("Media: decoding response body failed. Error: %s", err)
	}
	return TwitterResponse.MediaId, nil
}

// TwitterUploadMedia pipes the image resource(binary) to twitter media endpoint through multipart post,
// and returns the media_id_string if succeeds.
var UploadMedia = func(mediaUrl string, accessToken oauth.AccessToken) (mediaId string, err error) {
	mediaResponseBody, err := fetchMedia(mediaUrl)
	if err != nil {
		return "", err
	}
	defer mediaResponseBody.Close()
	mediaId, err = postMedia(mediaResponseBody, &accessToken)
	return mediaId, err
}

// TwitterStatusesUpdate post a new tweet to twitter,
// and return the tweet id(id_str) if succeeds.
// only handle this separately, others just throw errors as they are:
// Response Code: 403
// Response Body: {"errors":[{"code":186,"message":"Status is over 140 characters."}]}
var UpdateStatuses = func(status string, media_ids string, accessToken oauth.AccessToken) (tweetId string, err error) {
	params := map[string]string{
		"trim_user": "true",
		"status":    status,
	}
	// tweet with image is handled here
	if media_ids != "" {
		params["media_ids"] = media_ids
	}
	response, err := Consumer.Post(tweetEndpoint, params, &accessToken)
	if err != nil {
		if response.StatusCode == 403 {
			if strings.Contains(err.Error(), `"message":"Status is over 140 characters."`) {
				err = errors.New("Status is over 140 characters")
				//Twitter response
				// Get https://api.twitter.com/1.1/statuses/update.json returned status 403, {"errors":[{"code":186,"message":"Status is over 140 characters."}]}
			}
		}
		return "", err
	}
	// Here status code is 200, parse the response
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	var TwitterResponse struct {
		IdStr string `json:"id_str"`
	}
	decoder.Decode(&TwitterResponse)
	if err != nil {
		return "", fmt.Errorf("Statuses: decoding response body failed. %s", err)
	}
	return TwitterResponse.IdStr, nil
}
