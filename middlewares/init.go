package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(ctx *gin.Context) {
	start := time.Now()
	ctx.Set("username", "zzlion")
	ctx.Next()
	fmt.Println(ctx.Request.URL, time.Now().Sub(start).Nanoseconds())
}

func GoRoutineMiddleware(ctx *gin.Context) {
	cCp := ctx.Copy()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Done in path" + cCp.Request.URL.Path)
	}()
}
