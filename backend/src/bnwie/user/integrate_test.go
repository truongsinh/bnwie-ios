package user

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func TestUserRegAndLogin1(t *testing.T) {
	RegisterTestingT(t)

	t.Log(ts.URL + "/user/authenticate")
	res, err := http.Post(ts.URL+"/user/authenticate", "application/json", strings.NewReader(`{"FbsId": "123"}`))
	Expect(err).To(BeNil())
	Expect(res.StatusCode).To(Equal(400))
	body, err := ioutil.ReadAll(res.Body)
	Expect(err).To(BeNil())
	Expect(string(body)).To(MatchJSON(`{"Code":-1,"Error":"\"FbId\" required"}`))
}

func TestUserRegAndLogin2(t *testing.T) {
	RegisterTestingT(t)

	res, err := http.Post(ts.URL+"/user/authenticate", "application/json", strings.NewReader(`
		{
			"SocnetType": "facebook",
			"SocnetId": "123",
			"SocnetToken": "`+fbToken+`"
		}
		`))
	_, _ = res, err
	Expect(err).To(BeNil())
	Expect(res.StatusCode).To(Equal(200))
	body, err := ioutil.ReadAll(res.Body)
	Expect(err).To(BeNil())
	Expect(string(body)).To(MatchJSON(`{"Socnet":"facebook","Id":"1649497369","AvatarUrl":"https://scontent.xx.fbcdn.net/v/t1.0-1/p200x200/13466024_10208393622192018_8212025284330553497_n.jpg?oh=22a0f8bb503ceb297f9dcf5fb1b86e32\u0026oe=5834676D","ProfileUrl":"https://www.facebook.com/app_scoped_user_id/1649497369/","FullName":"TruongSinh Tran-Nguyen","Email":"i@truongsinh.pro","FirstName":"TruongSinh","LastName":"Tran-Nguyen","AgeRange":{"Min":21,"Max":0},"Gender":"male","Locale":"en_US","Timezone":3,"UpdatedTime":"2016-06-21T14:35:52+0000","Verified":true}`))
}
