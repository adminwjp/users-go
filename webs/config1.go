
package web

import (
"github.com/adminwjp/infrastructure-go/dtos"
"github.com/adminwjp/infrastructure-go/webs"
"github.com/adminwjp/users-go/models"
service "github.com/adminwjp/users-go/services"
"reflect"
)

type Config1Ctrl struct {
	Service func()service.ConfigService
}


func (ctrl *Config1Ctrl)M(httpWeb webs.HttpWeb){
	res:=dtos.ResponseDataDto{Status: true,Code:200,Msg:"success ",Data:&models.RpcModel{}}
	httpWeb.Response(200,res)
}
/*
	添加
*/
func (ctrl *Config1Ctrl)Add(httpWeb webs.HttpWeb){
	ctrl.Save(httpWeb,true)
}

/*
	修改
*/
func (ctrl *Config1Ctrl)Update(httpWeb webs.HttpWeb){
	ctrl.Save(httpWeb,false)
}
func (ctrl *Config1Ctrl)Save(httpWeb webs.HttpWeb,add bool){
	var m1 models.ConfigModel
	err:=httpWeb.ShouldBind(&m1)
	if err!=nil||reflect.DeepEqual(&m1,&models.ConfigModel{}){
		res:=dtos.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true, func(smsService service.ConfigService) {
		var r int
		n:="add"
		if add{
			r,_=smsService.Add(&m1)
		}else{
			r,_=smsService.Update(&m1)
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
func (ctrl *Config1Ctrl)Delete(httpWeb webs.HttpWeb){
	id:=httpWeb.GetPathInt64(":id")
	if id<1{
		res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "id error"}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true, func(smsService service.ConfigService) {
		if r,_:=smsService.Delete(id);r<1{
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
func (ctrl *Config1Ctrl)DeleteBatch(httpWeb webs.HttpWeb){
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
	ctrl.Action(httpWeb,true, func(smsService service.ConfigService) {
		r:=0
		if c1==webs.Xml{
			ids:=make([]int64,len(*&m1.Ids))
			for i := 0; i < len(*&m1.Ids); i++ {
				ids[i]=*&m1.Ids[i].Id
			}
			r,_=smsService.DeleteBatch(ids)
		}else{
			r,_=smsService.DeleteBatch(*&m.Ids)
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

func (ctrl *Config1Ctrl) List(httpWeb webs.HttpWeb){

	var page dtos.PageDto
	_=httpWeb.ShouldBindUri(&page)

	ctrl.Action(httpWeb,false, func(smsService service.ConfigService) {
		users,count,err1:=smsService.List(*&page.Page,*&page.Size)
		if err1!=nil{
			res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "list fail"}
			httpWeb.Response(200,res)
		}else{
			List(httpWeb,users,page.Page,page.Size,count)
		}
	})

}

func (ctrl *Config1Ctrl) Action(httpWeb webs.HttpWeb,write bool,
	method func( smsService service.ConfigService)){
	service:=ctrl.Service()
	if write{
		service.GetTranction().Begin()
	}

	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()
	//null
	method(service)
}
