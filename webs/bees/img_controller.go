package web_bee_controller

import (
	util "github.com/adminwjp/infrastructure-go/utils"
	web "github.com/adminwjp/users-go/webs"
	bee "github.com/beego/beego/v2/server/web"
)

type ImgController struct {
	bee.Controller
	web.ImgCtrl
}
func (ctrl *ImgController) Get(){
	img:=ctrl.GetString("img","")
	ctrl.ImgCtrl.Get(img,ctrl.Ctx.Request,ctrl.Ctx.ResponseWriter)
}

func (ctrl *ImgController) Upload(){
	id:=ctrl.GetString("id","")
	_,f,_ := ctrl.GetFile("pic")
	ctrl.SaveToFile(f.Filename,"static/imgs/"+id+util.FileUtil.GetFileExtension(f.Filename))
	ctrl.Data["json"]=map[string]interface{}{
		"status": true, "code": 200, "msg": "success",
	}
	ctrl.ServeJSON()
}