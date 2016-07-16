package facebook

import (
	"encoding/json"
	"fmt"
	"github.com/Smarp/socnet"
	"net/http"
)

type profileStruct struct {
	Socnet string
	Id     string `json:"id"`
	Link   string `json:"link"`

	Picture avatarStruct `json:"picture"`

	Email     string `json:"email"`
	FullName  string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	AgeRange struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"age_range"`
	Gender      string `json:"gender"`
	Locale      string `json:"locale"`
	Timezone    int    `json:"timezone"`
	UpdatedTime string `json:"updated_time"`
	Verified    bool   `json:"verified"`
}

func Profile(c *http.Client, accessToken string) (*socnet.Profile, error) {
	r, err := BatchParse(BatchRequest(c, accessToken, []*batchRequestStruct{myProfileRequest, myAvatarRequest}))
	if err != nil {
		return nil, err
	}
	p, err := ProfileParse(r[0])
	if err != nil {
		return nil, err
	}
	if a := avatarParse(r[1]); a != "" {
		p.AvatarUrl = a
	}
	return p, nil
}

const (
	endpoint        = "graph.facebook.com"
	graphApiVersion = "v2.7"
)

var myProfileRequest = newBatchRequest(
	"GET",
	"me",
	map[string]string{"fields": "name,email,first_name,last_name,age_range,link,picture,gender,locale,timezone,updated_time,verified"},
)

type avatarStruct struct {
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

var myAvatarRequest = newBatchRequest(
	"GET",
	"me/picture",
	map[string]string{"type": "large", "redirect": "false"},
)

func avatarParse(resp batchResponseStruct) string {

	if resp.StatusCode != http.StatusOK {
		// @todo
		fmt.Printf("avatar respose code %d", resp.StatusCode)
		return ""
	}
	data := avatarStruct{}
	json.Unmarshal([]byte(resp.Body), &data)
	return data.Data.Url
}

func ProfileParse(resp batchResponseStruct) (profile *socnet.Profile, err error) {
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Err Facebook Graph Profile statusCode %d %s", resp.StatusCode, resp.Body)
		return nil, err
	}
	profileRaw := profileStruct{}
	json.Unmarshal([]byte(resp.Body), &profileRaw)
	profile = &socnet.Profile{
		"facebook",
		profileRaw.Id,
		profileRaw.Picture.Data.Url,
		profileRaw.Link,
		profileRaw.FullName,
		profileRaw.Email,
		profileRaw.FirstName,
		profileRaw.LastName,
		struct {
			Min int
			Max int
		}{
			profileRaw.AgeRange.Min,
			profileRaw.AgeRange.Max,
		},
		profileRaw.Gender,
		profileRaw.Locale,
		profileRaw.Timezone,
		profileRaw.UpdatedTime,
		profileRaw.Verified,
	}
	return profile, err
}

func (*Socnet) MyProfile(accessToken string) (profile *socnet.Profile, err error) {
	return Profile(http.DefaultClient, accessToken)
}
