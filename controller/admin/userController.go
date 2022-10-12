package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go_gin/controller"
	"github.com/go_gin/models"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type UserController struct {
	controller.BaseController
}

func (con UserController) UserIndex(c *gin.Context) {
	user := []models.User{}
	// models.DB.Find(&user)
	models.DB.Where("id > 1").Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	// c.String(http.StatusOK, "admin 首页")
}

func (con UserController) UserList(c *gin.Context) {
	session := sessions.Default(c)
	cookie, err := c.Cookie("username")
	if err != nil {
		cookie = ""
	}
	// maxAge -1 删除cookie
	c.SetCookie("username", "", -1, "/", "localhost", false, true)

	value, ok := c.Get("username")
	if !ok {
		value = "无值"
	}
	v, ok := value.(string)
	if !ok {
		v = ""
	}
	fmt.Println("通过ctx传值" + v)
	// con.Success(c)
	c.String(http.StatusOK, "cookie:%s, session:%s", cookie, session.Get("username"))
}

func (con UserController) UserAdd(c *gin.Context) {
	user := models.User{
		Username: "zzlion",
		Age:      18,
		Email:    "163.com",
	}
	models.DB.Omit("AddTime", "UpdateTime").Create(&user)
	c.String(200, "用户列表-add")
}

func (con UserController) UserEdit(c *gin.Context) {
	// user := models.User{Id: 1}
	// models.DB.Find(&user)
	// user.Age = 16
	// models.DB.Save(&user)

	user := models.User{}
	models.DB.Model(&user).Where("id=?", 1).Update("age", 13)
	c.String(200, "用户列表-edit")
}

func (conn UserController) UserDelete(c *gin.Context) {
	// user := models.User{Id: 1}
	// models.DB.Find(&user)
	// models.DB.Delete(&user)

	user := models.User{}
	models.DB.Where("username = ?", "world").Delete(&user)
	c.String(http.StatusOK, "删除用户:%s", user.Username)
}

func (con UserController) UploadFilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/file.html", gin.H{})
}

func (con UserController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("face")
	allowMap := map[string]bool{".jpg": true, ".png": true, ".gif": true, ".jpeg": true}
	if err == nil {
		extName := path.Ext(file.Filename)
		_, ok := allowMap[extName]
		if ok {
			day := time.Now().Format("20060102")
			dir := "./static/upload/" + day
			err := os.MkdirAll(dir, 0766)
			if err != nil {
				c.String(http.StatusOK, "创建目录失败")
				return
			}
			timestamp := time.Now().Unix()
			fileName := strconv.FormatInt(timestamp, 10) + extName
			dst := path.Join(dir, fileName)
			fmt.Println(dst)
			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				c.String(http.StatusOK, "保存图片失败"+err.Error())
				return
			}
			c.String(http.StatusOK, "图片保存成功")
		} else {
			con.Error(c)
		}

	}
}

func (con UserController) UploadFilePageMulti(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/filemulti.html", gin.H{})
}

func (con UserController) UploadFileMulti(c *gin.Context) {
	file1, _ := c.FormFile("face1")
	log.Println(file1.Filename)

	dst := path.Join("./static/upload", file1.Filename)
	// 上传文件到指定目录
	err := c.SaveUploadedFile(file1, dst)

	file2, _ := c.FormFile("face2")
	log.Println(file2.Filename)

	dst = path.Join("./static/upload", file2.Filename)
	// 上传文件到指定目录
	err = c.SaveUploadedFile(file2, dst)

	if err != nil {
		con.Error(c)
	} else {
		con.Success(c)
	}

}

func (con UserController) UploadFilePageMultiSame(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/filemultisame.html", gin.H{})
}

func (con UserController) UploadFileMultiSame(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	var ret []bool
	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			ret = append(ret, false)
		} else {
			ret = append(ret, true)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(files),
	})
}
