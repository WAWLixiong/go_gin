package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func DataRouter(router *gin.Engine) {
	dataRouters := router.Group("/data")
	{
		dataRouters.GET("/json", func(ctx *gin.Context) {
			// ctx.JSON(http.StatusOK, map[string]any{
			// 	"success": true,
			// 	"msg":     "你好",
			// })
			// ctx.JSON(http.StatusOK, gin.H{
			// 	"success": true,
			// 	"msg":     "你好",
			// })
			// ctx.JSON(http.StatusOK, &Student{
			// 	Name: "zzlion",
			// 	Age:  18,
			// })

			// 解决跨域问题时会用到
			ctx.JSONP(http.StatusOK, &Student{
				Name: "zzlionl",
				Age:  18,
			})
		})

		dataRouters.GET("/xml", func(ctx *gin.Context) {
			ctx.XML(http.StatusOK, gin.H{
				"succes": true,
				"msg":    "xml",
			})

		})

		dataRouters.GET("/html", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "default/index.html", gin.H{
				"title": "我是后台数据",
				"name":  "h5",
				"age":   45,
				"article": &Article{
					Title: "中国",
				},
				"numbers": []int{1, 2, 3, 4},
				"articles": []any{
					&Article{
						Title: "中国",
					},
					&Article{
						Title: "中国",
					},
				},
				"empty": []any{},
				"ts":    1665384550,
			})
		})
	}
}
