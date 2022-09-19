package web_gin_controller

import (
	util "github.com/adminwjp/infrastructure-go/utils"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
)

type ImgController struct {
	web.ImgCtrl
}
func (ctrl *ImgController) Get(c *gin.Context){
	img:=c.GetString("img")
	ctrl.ImgCtrl.Get(img,c.Request,c.Writer)
}

func (ctrl *ImgController) Upload(c *gin.Context){
	id:=c.GetString("id")
	file,_ := c.FormFile("pic")
	c.SaveUploadedFile(file,"static/imgs/"+id+util.FileUtil.GetFileExtension(file.Filename))
	c.JSON(200,gin.H{
		"status": true, "code": 200, "msg": "success",
	})
}