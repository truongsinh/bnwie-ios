package user

import (
	"bnwie/user/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserUpdateByFbIdView = func(c *gin.Context) {
	auth := new(model.ExtendedData)
	if err := c.BindJSON(auth); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("abc"))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": `ExtendedData required`,
			"Code":  -1,
			"Debug": err.Error(),
		})
		return
	}

	if err := userUpdateByFbIdCtrl(auth); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("abc"))
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": `sthwrong`,
			"Code":  -1,
			"Debug": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"Status": "Updated"})
	return
}

var userUpdateByFbIdCtrl = func(a *model.ExtendedData) error {
	// 2. Upsert into DB
	if err := model.UpdateServiceByIdModel(a); err != nil {
		return err
	}
	return nil
}
