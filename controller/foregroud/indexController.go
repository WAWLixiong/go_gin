package foregroud

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 3600 * 6})
	session.Set("username", "zzlion123")
	session.Save()

	// .it.com 多个二级域名共享 cookie
	// a.it.com b.it.com都可以共享
	c.SetCookie("username", "张三", 3600, ".it.com", "localhost", false, true)
	c.String(http.StatusOK, "首页:%v\n", "你好")
}
