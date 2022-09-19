package web

import (
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/models"
	service "github.com/adminwjp/users-go/services"
	"reflect"
)

type EmailCtrl struct {
	Service func()service.EmailService
}
func (ctrl *EmailCtrl)M(httpWeb webs.HttpWeb){
	res:=dtos.ResponseDataDto{Status: true,Code:200,Msg:"success ",Data:&models.EmailConfigModel{}}
	httpWeb.Response(200,res)
}
/*
	添加
*/
func (ctrl *EmailCtrl)Add(httpWeb webs.HttpWeb){
	ctrl.Save(httpWeb,true)
}

/*
	修改
*/
func (ctrl *EmailCtrl)Update(httpWeb webs.HttpWeb){
	ctrl.Save(httpWeb,false)
}
func (ctrl *EmailCtrl)Save(httpWeb webs.HttpWeb,add bool){
	var m1 models.EmailConfigModel
	err:=httpWeb.ShouldBind(&m1)
	//bind fail bug
	if err!=nil{
		res:=dtos.ResponseDataDto{Status: false,Code:400,Msg:"fail "+err.Error(),Data:&m1}
		httpWeb.Response(200,res)
		return
	}
	if err!=nil||reflect.DeepEqual(&m1,&models.EmailConfigModel{}){
		res:=dtos.ResponseDataDto{Status: false,Code:400,Msg:"fail",Data:&m1}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true,  func(emailService service.EmailService) {
		var r int
		n:="add"
		if add{
			r,_=emailService.Add(&m1)
		}else{
			r,_=emailService.Update(&m1)
			n="update"
		}
		if r<1{
			res:=dtos.ResponseDto{Status: false, Code: 400, Msg: n+" fail"}
			httpWeb.Response(200,res)
		}else{
			res:=dtos.ResponseDto{Status: true, Code: 200, Msg: n+" success"}
			httpWeb.Response(200,res)
		}
	})


}
/*
	删除
*/
func (ctrl *EmailCtrl)Delete(httpWeb webs.HttpWeb){
	//id:=c.GetInt64(":id")
	//p:= c.Request.URL.Path
	//p=strings.ToLower(p)
	//m:=regexp.MustCompile(p)
	//bs:=m.Find([]byte("/role/remove/(\\d+)"))
	//id,_:=strconv.ParseInt(string(bs),10,64)
	id:=httpWeb.GetPathInt64(":id")
	if id<1{
		res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "id error"}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true,  func(emailService service.EmailService) {
		if r,_:=emailService.Delete(id);r<1{
			res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "delete fail"}
			httpWeb.Response(200,res)
		}else{
			res:=dtos.ResponseDto{Status: true, Code: 200, Msg: "delete success"}
			httpWeb.Response(200,res)
		}
	})


}

/*
	删除
*/
func (ctrl *EmailCtrl)DeleteBatch(httpWeb webs.HttpWeb){
	var m dtos.DeleteDto
	var m1 dtos.DeleteXmlDto
	err:=httpWeb.ShouldBindIds(&m,&m1)
	c:=httpWeb.ContentType()
	c1:= webs.ParseContentType(c)
	if err!=nil||(c1==webs.Xml&&reflect.DeepEqual(*&m1,dtos.DeleteXmlDto{}))||
		(c1!=webs.Xml&&reflect.DeepEqual(*&m,dtos.DeleteDto{})){
		res:=dtos.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true,  func(emailService service.EmailService) {
		r:=0
		if c1==webs.Xml{
			ids:=make([]int64,len(*&m1.Ids))
			for i := 0; i < len(*&m1.Ids); i++ {
				ids[i]=*&m1.Ids[i].Id
			}
			r,_=emailService.DeleteBatch(ids)
		}else{
			r,_=emailService.DeleteBatch(*&m.Ids)
		}
		if r<1{
			res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "delete fail"}
			httpWeb.Response(200,res)
		}else{
			res:=dtos.ResponseDto{Status: true, Code: 200, Msg: "delete success"}
			httpWeb.Response(200,res)
		}
	})


}

func (ctrl *EmailCtrl) List(httpWeb webs.HttpWeb){

	var page dtos.PageDto
	_=httpWeb.ShouldBindUri(&page)

	ctrl.Action(httpWeb,false, func(emailService service.EmailService) {
		users,count,err1:=emailService.List(*&page.Page,*&page.Size)
		if err1!=nil{
			res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
			httpWeb.Response(200,res)
		}else{
			List(httpWeb,users,page.Page,page.Size,count)
		}
	})
}
func (ctrl *EmailCtrl) Action(httpWeb webs.HttpWeb,write bool,
	method func(emailService service.EmailService)){
	service:=ctrl.Service()
	if write{
		service.GetTranction().Begin()
	}
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	method(service)
}

func ListByPage(httpWeb webs.HttpWeb,method func()(error,int,int,interface{},int64)){

	err,page,size,data,records:=method()
	if err!=nil{
		res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
		httpWeb.Response(200,res)
		return
	}else{
		List(httpWeb,data,page,size,records)
	}
}