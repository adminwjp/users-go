package web_http_controller

import (
	util "github.com/adminwjp/infrastructure-go/utils"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/infrastructure-go/webs/https"
	web "github.com/adminwjp/users-go/webs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type ImgController struct {
	web.ImgCtrl
}
func (ctrl *ImgController) Get(writer http.ResponseWriter,request *http.Request){
	img:=request.URL.Query().Get("img")
	ctrl.ImgCtrl.Get(img,request,writer)
}

func (ctrl *ImgController) Upload(writer http.ResponseWriter,request *http.Request){

	id:=request.URL.Query().Get("id")

	file,f,_ := request.FormFile("pic")
	bs,_:=ioutil.ReadAll(file)
	ioutil.WriteFile("static/imgs/"+id+util.FileUtil.GetFileExtension(f.Filename),bs,
		os.ModeType)
	https.OutputXmlOrJson(writer,webs.Json,gin.H{
		"status": true, "code": 200, "msg": "success",
	})
}