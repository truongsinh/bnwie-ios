package user

import (
	cv "bnwie/user/cv"
	model "bnwie/user/model"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// 1st user
// login with FB
// about me from fb

// chọn make-up artist
// nghệ sĩ trang điểm
// hair
// face
// nail
// body
// make-up consumer
// người dùng

func RegisterHandler(g *gin.RouterGroup, db *sql.DB) {
	model.PrepareAll(db)
	g.POST("/authenticate", cv.UserAuthenticateView)
	g.POST("/", cv.UserUpdateByFbIdView)
}
