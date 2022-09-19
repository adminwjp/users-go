package web

import (
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/adminwjp/users-go/models"
	service "github.com/adminwjp/users-go/services"
	"reflect"
)

type PayCtrl struct {
	Service func()service.PaySecrtService
}
func (ctrl *PayCtrl)M(httpWeb webs.HttpWeb){
	res:=dtos.ResponseDataDto{Status: true,Code:200,Msg:"success ",Data:&models.PaySecrtConfigModel{}}
	httpWeb.Response(200,res)
}
/*
	添加
*/
func (ctrl *PayCtrl)Add(httpWeb webs.HttpWeb){
	ctrl.Save(httpWeb,true)
}

/*
	修改
*/
func (ctrl *PayCtrl)Update(httpWeb webs.HttpWeb){
	ctrl.Save(httpWeb,false)
}
func (ctrl *PayCtrl)Save(httpWeb webs.HttpWeb,add bool){
	var m1 models.PaySecrtConfigModel
	err:=httpWeb.ShouldBind(&m1)
	if err!=nil||reflect.DeepEqual(&m1,&models.PaySecrtConfigModel{}){
		res:=dtos.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true, func(paySecrtService service.PaySecrtService) {
		var r int
		if add{
			r,_=paySecrtService.Add(&m1)
		}else{
			r,_=paySecrtService.Update(&m1)
		}
		if r<1{
			res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "update fail"}
			httpWeb.Response(200,res)
		}else{
			res:=dtos.ResponseDto{Status: true, Code: 200, Msg: "update success"}
			httpWeb.Response(200,res)
		}
	})

}
/*
	删除
*/
func (ctrl *PayCtrl)Delete(httpWeb webs.HttpWeb){
	id:=httpWeb.GetPathInt64(":id")
	if id<1{
		res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "id error"}
		httpWeb.Response(200,res)
		return
	}
	ctrl.Action(httpWeb,true, func(paySecrtService service.PaySecrtService) {
		if r,_:=paySecrtService.Delete(id);r<1{
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
func (ctrl *PayCtrl)DeleteBatch(httpWeb webs.HttpWeb){
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
	ctrl.Action(httpWeb,true, func(secrtService service.PaySecrtService) {
		r:=0
		if c1==webs.Xml{
			ids:=make([]int64,len(*&m1.Ids))
			for i := 0; i < len(*&m1.Ids); i++ {
				ids[i]=*&m1.Ids[i].Id
			}
			r,_=secrtService.DeleteBatch(ids)
		}else{
			r,_=secrtService.DeleteBatch(*&m.Ids)
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

func (ctrl *PayCtrl) List(httpWeb webs.HttpWeb){

	var page dtos.PageDto
	_=httpWeb.ShouldBindUri(&page)

	service:=ctrl.Service()
	//service.GetTranction().Begin()
	defer func() {
		service.GetTranction().Commit()
		service.Clean()
		service=nil
	}()

}

func (ctrl *PayCtrl) Action(httpWeb webs.HttpWeb,write bool,
	method func( service.PaySecrtService)){
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
