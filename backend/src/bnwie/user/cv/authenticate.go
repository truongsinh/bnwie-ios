package user

import (
	"bnwie/user/model"
	"errors"
	"net/http"

	"github.com/Smarp/socnet"
	"github.com/Smarp/socnet/facebook"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type authStruct struct {
	// FB
	SocnetType  string `binding:"required"`
	SocnetId    string `binding:"required"`
	SocnetToken string `binding:"required"`
}

var UserAuthenticateView = func(c *gin.Context) {
	auth := new(authStruct)
	if err := c.BindJSON(auth); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("abc"))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": `"FbId" required`,
			"Code":  -1,
		})
		return
	}

	if resp, err := userAuthenticateCtrl(auth); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("abc"))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": `sthwrong`,
			"Code":  -1,
			"Debug": err.Error(),
		})
		return
	} else {
		c.JSON(200, resp)
		return
	}
}
var fbOAuthConsumer = &facebook.Socnet{}
var userAuthenticateCtrl = func(a *authStruct) (*socnet.Profile, error) {
	// 1. Get Profile with provided FB token
	p, err := fbOAuthConsumer.MyProfile(a.SocnetToken)

	u := &model.User{
		Profile:      p,
		ExtendedData: &model.ExtendedData{},
	}
	if err != nil {
		return nil, err
	}

	// 2. Upsert into DB
	if err := model.UpsertUserModel(u); err != nil {
		return nil, err
	}

	return p, nil
}
