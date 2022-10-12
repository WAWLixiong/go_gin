package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_gin/controller/admin"
)

func ArticleRouter(ctx *gin.Engine) {
	r := ctx.Group("/article")
	r.GET("/news", admin.ArticleController{}.UserList)
	r.GET("/news/add", admin.ArticleController{}.UserAdd)
	r.GET("/news/edit", admin.ArticleController{}.UserEdit)

}
