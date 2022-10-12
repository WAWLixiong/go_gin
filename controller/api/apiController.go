package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go_gin/controller"
)

type ApiController struct {
	controller.BaseController
}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "接口")
}

func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "用户列表")
}

func (con ApiController) PList(c *gin.Context) {
	c.String(200, "新闻列表")
}
