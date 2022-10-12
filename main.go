package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go_gin/models"
	"github.com/go_gin/routers"
	"html/template"
)

// func UnixToTime(timestamp int) string {
// 	t := time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
// 	return t
// }

func Println(str1, str2 string) string {
	return str1 + "---" + str2
}

func initMiddlewareOne(ctx *gin.Context) {
	fmt.Println("global mid 1 in")
	ctx.Next()
	fmt.Println("global mid 1 out")
}

func initMiddlewareTwo(ctx *gin.Context) {
	fmt.Println("global mid 2 in")
	ctx.Next()
	fmt.Println("global mid 2 out")
}

func main() {
	r := gin.Default()
	// 配置函数放在加载模板之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Println":    Println, // 多个参数
	})
	r.LoadHTMLGlob("templates/**/*") // 放在配置路由之前
	r.Static("/static", "./static")

	// session存储引擎
	// cookie
	// store := cookie.NewStore([]byte("secret_key"))
	// redis
	store, err := redis.NewStore(10, "tcp", "localhost:49153", "redispw", []byte("mysession"))
	if err != nil {
		fmt.Println(err)
		return
	}

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(initMiddlewareOne, initMiddlewareTwo, sessions.Sessions("mysession", store))

	// 抽离到独立模块
	// defaultRouters := r.Group("/")
	// {
	// 	defaultRouters.GET("/", func(context *gin.Context) {
	// 		context.String(http.StatusOK, "值是:%v\n", "你好")
	// 	})
	// }

	routers.DefaultRouter(r)
	routers.DataRouter(r)
	routers.UserRouter(r)
	routers.ParamRouter(r)
	routers.ApiRouter(r)

	r.Run()
}
