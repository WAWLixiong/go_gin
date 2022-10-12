package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_gin/controller/api"
)

func ApiRouter(ctx *gin.Engine) {
	r := ctx.Group("api")
	r.GET("/", api.ApiController{}.Index)
	r.GET("/userlist", api.ApiController{}.UserList)
	r.GET("/plist", api.ApiController{}.PList)
}
