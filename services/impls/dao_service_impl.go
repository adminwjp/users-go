package service_impl

import (
	config_es "github.com/adminwjp/infrastructure-go/configs/ess"
	config_gorm "github.com/adminwjp/infrastructure-go/configs/gorms"
	config_gorm_jinzhu "github.com/adminwjp/infrastructure-go/configs/gorms/jinzhus"
	config_mong "github.com/adminwjp/infrastructure-go/configs/mongs"
	data "github.com/adminwjp/infrastructure-go/datas"
	data_db_bee "github.com/adminwjp/infrastructure-go/datas/dbs/bees"
	data_db_gorm "github.com/adminwjp/infrastructure-go/datas/dbs/groms"
	data_db_gorm_jinzhu "github.com/adminwjp/infrastructure-go/datas/dbs/groms/jinzhus"
	dao_gorm_impl "github.com/adminwjp/users-go/daos/impls/gorm"
	"log"

	"github.com/adminwjp/users-go/daos"
	dao_bee_orm_impl "github.com/adminwjp/users-go/daos/impls/bee_orm"
	dao_es_impl "github.com/adminwjp/users-go/daos/impls/es"
	dao_gorm_jinzhu_impl "github.com/adminwjp/users-go/daos/impls/gorm/jinzhu"
	dao_mong_impl "github.com/adminwjp/users-go/daos/impls/mong"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/models"
)

type DaoServiceImpl struct {
	Data func()daos.Dao
	//es mong
	enableEs bool
	enableMong bool
	//mong es 2 select 1
	//go 自动回收 不能 用 最好全局的不然坑多 局部的 私有的不行 cruds Cruds 也不行
	Cruds map[string]daos.CrudDao
	//no used
	crudEss map[string]daos.CrudDao
	crudMongs map[string]daos.CrudDao
}


func (dao *DaoServiceImpl)UpdateDataUser(serivce *UserServiceImpl)  {
	switch datas.GlobalConfig.DataFlag {
	case data.DataDb:
		switch  datas.GlobalConfig.OrmFlag {
		case data.DbOrmJinzhuGorm:
			db:=config_gorm_jinzhu.GormConfigInstance.Db
			trans:=&data_db_gorm_jinzhu.TranManager{Db:db }
			dao.updateUser(serivce,trans)
			break
		case data.DbOrmGormio:
			db:=config_gorm.GormConfigInstance.Db
			trans:=&data_db_gorm.TranManager{Db:db }
			dao.updateUser(serivce,trans)
			break
		case data.DbOrmBee:
			trans:=&data_db_bee.TranManager{}
			dao.updateUser(serivce,trans)
			break
		default:break

		}
		break
	case data.DataEs, data.DataMong:
		dao.Init(nil)

		daos.UserDaoInstance.Data=dao.Data
		daos.UserDaoInstance.TranManager=data.TranManagerEmpty
		serivce.Dao=daos.UserDaoInstance
		serivce.BaseDao=daos.UserDaoInstance

		break
	default:
		break
	}


}

func (dao *DaoServiceImpl)updateUser(serivce *UserServiceImpl,tans daos.TranDao)  {
	var dao1 =&daos.UserDaoImpl{}
	dao1.SetAdmin(false)
	dao1.SetUserSql()
	dao1.TranManager=tans

	dao.Init(dao1.TranManager)
	dao1.Data=dao.Data

	//var baseUserDao dao.BaseUserDao=dao
	serivce.BaseDao=dao1
	//var adminDao dao.AdminDao=dao
	serivce.Dao=dao1


}
func (dao *DaoServiceImpl)UpdateDataAdmin(serivce *AdminServiceImpl)  {
	switch datas.GlobalConfig.DataFlag {
	case data.DataDb:
		switch  datas.GlobalConfig.OrmFlag {
		case data.DbOrmJinzhuGorm:
			db:=config_gorm_jinzhu.GormConfigInstance.Db
			trans:=&data_db_gorm_jinzhu.TranManager{Db:db }
			dao.updateAdmin(serivce,trans)
			break
		case data.DbOrmGormio:
			db:=config_gorm.GormConfigInstance.Db
			trans:=&data_db_gorm.TranManager{Db:db }
			dao.updateAdmin(serivce,trans)
			break
		case data.DbOrmBee:
			trans:=&data_db_bee.TranManager{}
			dao.updateAdmin(serivce,trans)
			break
		default:break

		}
		break
	case data.DataEs:
	case data.DataMong:
		serivce.Dao.Init()
		//dao.Init(nil)
		//daos.AdminDaoInstance.SetAdmin(true)
		//daos.AdminDaoInstance.SetUserSql()

		daos.AdminDaoInstance.Data=dao.Data
		daos.AdminDaoInstance.TranManager=data.TranManagerEmpty
		serivce.Dao=daos.AdminDaoInstance
		serivce.BaseDao=daos.AdminDaoInstance

		break
	default:
		break
	}


}

func (dao *DaoServiceImpl)updateAdmin(serivce *AdminServiceImpl,tans daos.TranDao)  {
	var dao1 =&daos.AdminDaoImpl{}
	dao1.SetAdmin(true)
	dao1.SetUserSql()
	dao1.TranManager=tans

	dao.Init(dao1.TranManager)
	dao1.Data=dao.Data

	//var baseUserDao dao.BaseUserDao=dao
	serivce.BaseDao=dao1
	//var adminDao dao.AdminDao=dao
	serivce.Dao=dao1


}
var RoleServiceInstance *RoleServiceImpl
func (dao *DaoServiceImpl)GetRole( )*RoleServiceImpl{
	switch datas.GlobalConfig.DataFlag {
	case data.DataEs,data.DataMong:
		_,e:=dao.Cruds["role"]
		if e{
			//pass
			RoleServiceInstance.Dao=dao.Cruds["role"]
			RoleServiceInstance.RoleDao=daos.RoleDaoInstance
			return RoleServiceInstance
		}
		break
	default:
		break
	}
	var serivce =&RoleServiceImpl{}
	RoleServiceInstance=serivce
	dao.CrudData("role",&models.RoleModel{},&serivce.ServiceImpl)


	//override
	d:=serivce.Dao.(*daos.CrudDaoImpl)
	dao1:=&daos.RoleDaoImpl{TranManager:serivce.TranManager }
	daos.RoleDaoInstance=dao1
	dao1.Data=d.Data
	dao1.Model(&models.RoleModel{})
	serivce.RoleDao=dao1
	serivce.Dao=dao1
	return serivce
}
var RpcServiceInstance *RpcServiceImpl
func (dao *DaoServiceImpl)GetRpc()*RpcServiceImpl{
	if dao.isExists("rpc"){
		ConfigServiceInstance.Dao=dao.Cruds["rpc"] //pass
		return RpcServiceInstance}
	var serivce =&RpcServiceImpl{}
	RpcServiceInstance=serivce
	dao.CrudData("rpc",&models.RpcModel{},&serivce.ServiceImpl)
	daos.RpcDaoInstance=serivce.Dao.(*daos.CrudDaoImpl)
	return serivce
}
var ConfigServiceInstance *ConfigServiceImpl
func (dao *DaoServiceImpl)GetConfig()*ConfigServiceImpl{
	if dao.isExists("config"){
		ConfigServiceInstance.Dao=dao.Cruds["config"] //pass
		return ConfigServiceInstance}
	var serivce =&ConfigServiceImpl{}
	ConfigServiceInstance=serivce
	dao.CrudData("config",&models.ConfigModel{},&serivce.ServiceImpl)
	daos.ConfigDaoInstance=serivce.Dao.(*daos.CrudDaoImpl)
	return serivce
}
var SmsServiceInstance *SmsServiceImpl
func (dao *DaoServiceImpl)GetSms( )*SmsServiceImpl{
	//https://blog.csdn.net/TurkeyCock/article/details/86739810
	//first pass after bug
	if dao.isExists("sms"){
		log.Printf("sms service create success,return single")
		//SmsServiceInstance.Dao=daos.SmsDaoInstance //pass
		SmsServiceInstance.Dao=dao.Cruds["sms"] //pass
		return SmsServiceInstance}
	var serivce =&SmsServiceImpl{}
	SmsServiceInstance=serivce
	dao.CrudData("sms",&models.SmsConfigModel{},&serivce.ServiceImpl)
	log.Printf("sms service create success")
	//daos.SmsDaoInstance=serivce.Dao.(*daos.CrudDaoImpl)//pass
	//serivce.ServiceImpl=*&serivce.ServiceImpl //err
	//SmsServiceInstance.Dao=serivce.Dao //err
	//SmsServiceInstance.Dao=*&serivce.ServiceImpl.Dao //err
	//SmsServiceInstance.Dao=serivce.ServiceImpl.Dao //err
	return serivce
}
func (dao *DaoServiceImpl) isExists(name string )(bool)  {
	switch datas.GlobalConfig.DataFlag {
	case data.DataEs,data.DataMong:
		_,e:=dao.Cruds[name]
		if e{
			return true
		}
		break
	default:
		break
	}
	return  false
}
var PayServiceInstance *PayServiceImpl
func (dao *DaoServiceImpl)GetPay() *PayServiceImpl{
	var serivce=PayServiceInstance
	if dao.isExists("pay"){
		PayServiceInstance.Dao=dao.Cruds["pay"] //pass
		return PayServiceInstance
	}else{
		serivce =&PayServiceImpl{}
		PayServiceInstance=serivce
	}
	dao.CrudData("pay",&models.PaySecrtConfigModel{},&serivce.ServiceImpl)
	return serivce
}
var EmailServiceInstance *EmailServiceImpl
func (dao *DaoServiceImpl)GetEmail()*EmailServiceImpl{
	if dao.isExists("email"){
		EmailServiceInstance.Dao=dao.Cruds["email"] //pass
		return EmailServiceInstance}
	var serivce =&EmailServiceImpl{}
	EmailServiceInstance=serivce
	dao.CrudData("email",&models.EmailConfigModel{},&serivce.ServiceImpl)
	return serivce
}
func (dao *DaoServiceImpl)crud(m daos.MDao,serivce *ServiceImpl){

	dao.Init(serivce.TranManager)

	dao1:=&daos.CrudDaoImpl{}
	dao1.Model(m)
	serivce.Dao=dao1
	if serivce.Dao==nil{
		log.Println("crud service Dao is nil ")
	}
	dao1.Data=dao.Data

}
func (dao *DaoServiceImpl)CrudData(name string,m daos.MDao,serivce *ServiceImpl){
	switch datas.GlobalConfig.DataFlag {
	case data.DataDb:
		switch  datas.GlobalConfig.OrmFlag {
		case data.DbOrmJinzhuGorm:
			db:=config_gorm_jinzhu.GormConfigInstance.Db
			serivce.TranManager=&data_db_gorm_jinzhu.TranManager{Db:db }
			dao.crud(m,serivce)
			break
		case data.DbOrmGormio:
			db:=config_gorm.GormConfigInstance.Db
			serivce.TranManager=&data_db_gorm.TranManager{Db:db }
			dao.crud(m,serivce)
			break
		case data.DbOrmBee:
			serivce.TranManager=&data_db_bee.TranManager{}
			dao.crud(m,serivce)
			break
		default:break

		}
		break
	case data.DataEs, data.DataMong:
		_,v:=dao.Cruds[name]
		if v{
			return
		}

		serivce.TranManager=data.TranManagerEmpty
		dao.crud(m,serivce)
		dao.Cruds[name]=serivce.Dao
		break
	default:
		break
	}

}

func (dao *DaoServiceImpl)Init(tans daos.TranDao)  {

	switch datas.GlobalConfig.DataFlag {
	case data.DataDb:
		switch  datas.GlobalConfig.OrmFlag {
		case data.DbOrmJinzhuGorm:
			dao.Data= func() daos.Dao {
				tr:=tans.(*data_db_gorm_jinzhu.TranManager)
				da:= &dao_gorm_jinzhu_impl.GormDaoImpl{ TranManager:tr }
				return da
			}
			break
		case data.DbOrmGormio:
			dao.Data= func() daos.Dao {
				tr:=tans.(*data_db_gorm.TranManager)
				da:= &dao_gorm_impl.GormDaoImpl{ TranManager:tr }

				return da
			}
			break
		case data.DbOrmBee:
			dao.Data= func() daos.Dao {
				tr:=tans.(*data_db_bee.TranManager)
				da:= &dao_bee_orm_impl.BeeOrmDaoImpl{ TranManager:tr }
				return da
			}
			break
		default:break

		}
		break
	case data.DataEs:
		if dao.enableEs{
			return
		}
		dao.enableEs=true
		dao.Data= func() daos.Dao {
			log.Println("es Client init")
			if config_es.EsConfigInstance.Client==nil{
				log.Println("es Client init fail,is nil ")
			}
			da:= &dao_es_impl.EsDaoImpl{Client: config_es.EsConfigInstance.Client}
			//da.DaoImpl=daos.DaoImpl{}
			return da
		}
		break
	case data.DataMong:
		if dao.enableMong{
			return
		}
		dao.enableMong=true
		dao.Data= func() daos.Dao {
			log.Println("mong session init")
			if config_mong.MongConfigInstance.Session==nil{
				log.Println("mong session init fail,is nil ")
			}
			da:= &dao_mong_impl.MongDaoImpl{ Session: config_mong.MongConfigInstance.Session}
			//da.DaoImpl=daos.DaoImpl{}
			return da
		}
		break
	default:
		break
	}
}
