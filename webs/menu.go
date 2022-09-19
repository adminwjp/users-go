package web

import (
	"encoding/json"
	"github.com/adminwjp/infrastructure-go/webs"
	"log"
	"os"
	"strings"
)

type MenuCtrl struct {

}

func (*MenuCtrl)Load(httpWeb webs.HttpWeb)  {
	var dir,_=os.Getwd()
	log.Println(dir)
	var temp=strings.Replace(dir,"\\","/",0)
	bs,_:=os.ReadFile(temp+"/json/menu.json")
	log.Println(string(bs))
	var data []map[string]interface{}
	er:=json.Unmarshal(bs,&data)
	if er!=nil{
		log.Println(er.Error())
	}
	httpWeb.Json(200,map[string]interface{}{
		"status": true, "code": 200, "msg": "success","data":&data,
	})
}