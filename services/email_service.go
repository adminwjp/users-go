package service


import (
	"github.com/adminwjp/users-go/daos"
	"github.com/adminwjp/users-go/models"
)

//邮箱配置服务接口
type EmailService interface {
	Clean()
	GetTranction()daos.TranDao
	/**
	添加
	*/
	Add(model *models.EmailConfigModel)(int,error)

	/**
	修改
	*/
	Update(model *models.EmailConfigModel)(int,error)

	/**
	删除
	*/
	Delete(id int64)(int,error)

	/**
	删除
	*/
	DeleteBatch(ids []int64)(int,error)

	/**
	删除
	*/
	//DeleteByEmail(email string)int
	List1()([]models.EmailConfigModel,int64,error)
	List(page int,size int)([]models.EmailConfigModel,int64,error)
}
