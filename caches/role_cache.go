package caches

import "github.com/adminwjp/users-go/models"

//角色接口
type RoleCache interface {
	/*
		更新
	*/
	Save(role *models.RoleModel)(int,error)

	/*
		查询
	*/
	Get()([]models.RoleModel,error)


}
