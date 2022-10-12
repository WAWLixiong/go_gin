package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_gin/controller/admin"
	"github.com/go_gin/middlewares"
	"time"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func initMiddlewareOne(c *gin.Context) {
	fmt.Println("mid1 in")
	start := time.Now().UnixNano()
	c.Next()
	// c.Abort() // 中止后续的handler
	span := time.Now().UnixNano() - start
	fmt.Println("时间消耗:", span, "nano")
	fmt.Println("mid1 out")
}

func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("mid2 in")
	start := time.Now().UnixNano()
	c.Next()
	// c.Abort() // 中止后续的handler
	span := time.Now().UnixNano() - start
	fmt.Println("时间消耗:", span, "nano")
	fmt.Println("mid2 out")
}

func UserRouter(router *gin.Engine) {
	r := router.Group("/admin")

	r.Use(middlewares.InitMiddleware, middlewares.GoRoutineMiddleware)

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
	// })
	//
	// r.POST("/doAddUser", func(ctx *gin.Context) {
	// 	username := ctx.PostForm("username")
	// 	password := ctx.PostForm("password")
	// 	age := ctx.DefaultPostForm("age", "18")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 		"age":      age,
	// 	})
	// })
	//
	// // GET/POST数据保定到结构体上
	// r.GET("/getuser", func(ctx *gin.Context) {
	// 	user := &UserInfo{}
	// 	if err := ctx.ShouldBind(user); err == nil {
	// 		ctx.JSON(http.StatusOK, user)
	// 	} else {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{
	// 			"err": err.Error(),
	// 		})
	// 	}
	// })
	//
	// r.POST("/postuser", func(ctx *gin.Context) {
	// 	user := &UserInfo{}
	// 	if err := ctx.ShouldBind(user); err == nil {
	// 		ctx.JSON(http.StatusOK, user)
	// 	} else {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{
	// 			"err": err.Error(),
	// 		})
	// 	}
	// })

	userController := admin.UserController{}

	r.GET("/", userController.UserIndex)
	r.GET("/user", initMiddlewareOne, initMiddlewareTwo, userController.UserList)
	r.GET("/user/add", initMiddlewareOne, initMiddlewareTwo, userController.UserAdd)
	r.GET("/user/delete", initMiddlewareOne, initMiddlewareTwo, userController.UserDelete)
	r.GET("/user/edit", initMiddlewareOne, initMiddlewareTwo, userController.UserEdit)
	r.GET("/user/file", userController.UploadFilePage)
	r.POST("/user/uploadFile", initMiddlewareOne, initMiddlewareTwo, userController.UploadFile)
	r.GET("/user/fileMulti", userController.UploadFilePageMulti)
	r.POST("/user/uploadFileMulti", userController.UploadFileMulti)
	r.GET("/user/fileMultiSame", userController.UploadFilePageMultiSame)
	r.POST("/user/uploadFileMultiSame", userController.UploadFileMultiSame)
}
