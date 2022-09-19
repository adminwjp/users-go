##### go example
```go
import (
	"github.com/adminwjp/infrastructure-go/data/mong"
	dtos "github.com/adminwjp/infrastructure-go/dto"
	"github.com/adminwjp/users-go/dto"
	"github.com/adminwjp/users-go/model"
	"log"
)

//管理员接口 mong 实现
type AdminDaoImpl struct {
	mong.MongData
}
/*
	根据手机号、邮箱、用户名登录
*/
func(admin *AdminDaoImpl) Login(user dto.UserInput)*model.AdminModel{
	var where map[string]interface{}
	log.Printf("admin mong account login: account %s ",user.Account)

	var adminModel *model.AdminModel
	if user.Flag==dtos.AccounTypeByEamil{
		where=map[string]interface{}{"email":user.Account,"pwd":user.Pwd}
	}else if user.Flag==dtos.AccounTypeByUsername{
		where=map[string]interface{}{"user_name":user.Account,"pwd":user.Pwd}
	}else{
		where=map[string]interface{}{"phone":user.Account,"pwd":user.Pwd}
	}
	err:=admin.Session.DB("users").C("admin").Find(where).One(&adminModel)
	if err!=nil{
		log.Printf("admin mong account login fail: account %s,err => %s",
			user.Account,err.Error())
	}
	return  adminModel
}
```