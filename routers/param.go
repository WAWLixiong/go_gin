package routers

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParamRouter(router *gin.Engine) {
	r := router.Group("/parameter")
	// params 传值
	r.GET("/get/params", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")
		ctx.JSONP(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	// post传值
	r.POST("/post/form", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		age := ctx.DefaultPostForm("age", "10")

		ctx.JSONP(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
		})

	})
	r.POST("/xml", func(ctx *gin.Context) {
		article := &Article{}
		data, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		} else {
			xml.Unmarshal(data, article)
			ctx.JSON(http.StatusOK, article)
		}
	})
	// 动态路由
	r.GET("/news/:newsid", func(ctx *gin.Context) {
		newsid := ctx.Param("newsid")
		ctx.JSON(http.StatusOK, gin.H{
			"newsid": newsid,
		})
	})
}
