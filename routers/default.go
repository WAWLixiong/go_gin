package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_gin/controller/foregroud"
)

func DefaultRouter(router *gin.Engine) {
	defaultRouters := router.Group("/")
	defaultRouters.GET("/", foregroud.IndexController{}.Index)
}
