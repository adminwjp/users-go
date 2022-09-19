package daos

var RoleDaoInstance=&RoleDaoImpl{}

//map 引用自动清空了
var SmsDaoInstance=&CrudDaoImpl{}
var PayDaoInstance=&CrudDaoImpl{}
var EmailDaoInstance=&CrudDaoImpl{}

var ConfigDaoInstance=&CrudDaoImpl{}
var RpcDaoInstance=&CrudDaoImpl{}

var AdminDaoInstance=&AdminDaoImpl{}
var UserDaoInstance=&UserDaoImpl{}
var MDb="samplesystem"

func init()  {
	AdminDaoInstance.SetAdmin(true)
	AdminDaoInstance.SetUserSql()

	UserDaoInstance.SetAdmin(false)
	UserDaoInstance.SetUserSql()
}