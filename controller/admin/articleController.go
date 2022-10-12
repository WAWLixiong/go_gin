package admin

import "github.com/gin-gonic/gin"

type ArticleController struct{}

func (con ArticleController) UserList(c *gin.Context) {
	c.String(200, "文章列表")
}

func (con ArticleController) UserAdd(c *gin.Context) {
	c.String(200, "文章列表-add")
}

func (con ArticleController) UserEdit(c *gin.Context) {
	c.String(200, "文章列表-edit")
}
