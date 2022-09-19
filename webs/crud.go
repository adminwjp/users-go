package web

import (
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	service "github.com/adminwjp/users-go/services"
	"reflect"
)

type CrudCtrl struct {
	Service func()service.CrudService
}
/*
	添加
*/
func (ctrl *CrudCtrl)Add(httpWeb webs.HttpWeb,m1 interface{}){
	err:=httpWeb.ShouldBind(m1)
	if err!=nil||reflect.DeepEqual(m1,m1){
		res:=dtos.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	service:=ctrl.Service()
	if r,_:=service.Add(&m1);r<1{
		res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "add fail"}
		httpWeb.Response(200,res)
		return
	}else{
		res:=dtos.ResponseDto{Status: true, Code: 200, Msg: "add success"}
		httpWeb.Response(200,res)
		return
	}
}

/*
	修改
*/
func (ctrl *CrudCtrl)Update(httpWeb webs.HttpWeb,m1 interface{}){
	err:=httpWeb.ShouldBind(m1)
	if err!=nil||reflect.DeepEqual(m1,m1){
		res:=dtos.ResponseDto{Status: false,Code:400,Msg:"fail"}
		httpWeb.Response(200,res)
		return
	}
	service:=ctrl.Service()
	if r,_:=service.Update(&m1);r<1{
		res:=dtos.ResponseDto{Status: false, Code: 400, Msg: "add fail"}
		httpWeb.Response(200,res)
		return
	}else{
		res:=dtos.ResponseDto{Status: true, Code: 200, Msg: "add success"}
		httpWeb.Response(200,res)
		return
	}
}

/*
	删除
*/
func (ctrl *CrudCtrl)Delete(httpWeb webs.HttpWeb){
	//id:=httpWeb.GetPathInt64(":id")

}

/*
	删除
*/
func (ctrl *CrudCtrl)DeleteBatch(httpWeb webs.HttpWeb){

}

func (ctrl *CrudCtrl) List(httpWeb webs.HttpWeb){

}
