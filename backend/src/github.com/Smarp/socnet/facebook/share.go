package facebook

import (
	"encoding/json"
	"fmt"
	"github.com/Smarp/socnet"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var Share = func(c *http.Client, fbId string, shareObj *socnet.Share) (id string, err error) {
	return ShareParse(c.Do(ShareRequest(fbId, shareObj)))
}

func ShareRequest(userId string, s *socnet.Share) (req *http.Request) {
	shareUrl := "https://graph.facebook.com/v2.3/" + userId + "/feed?"
	q := url.Values{}
	q.Set("format", "json")
	q.Set("message", s.Message)
	q.Set("picture", s.ImageUrl)
	q.Set("link", s.Link)
	q.Set("name", s.Title)
	if s.Signature != "" {
		q.Set("caption", s.Signature)
	} else {
		q.Set("caption", s.CanonicalUrl)
	}
	q.Set("description", s.Description)
	req, _ = http.NewRequest("POST", shareUrl+q.Encode(), nil)
	return req
}

func ShareParse(resp *http.Response, ierr error) (shareId string, err error) {
	if ierr != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		byteBody, errCode := ioutil.ReadAll(resp.Body)
		if errCode != nil {
			errCode = fmt.Errorf("facebook.ShareParse: status code not 200 but %d. Cannot read body: %s", resp.StatusCode, errCode)
			return "", errCode
		}
		errCode = fmt.Errorf("facebook.ShareParse: status code not 200 but %d. Response body: %s", resp.StatusCode, string(byteBody))
		return "", errCode
	}
	var resBody = struct {
		Id string `json:"id"`
	}{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&resBody)
	if err != nil {
		err = fmt.Errorf("facebook.ShareParse: cannot decode body: %s\n", err)
		return "", err
	}
	shareId = strings.Split(resBody.Id, "_")[1]
	return shareId, nil
}
