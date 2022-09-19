package daos

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/dtos"
)

/*
	根据手机号修改手机号
*/
func (dao *BaseUserDaoImpl)UpdatePhone(phone string,newPhone string)(int,error){
	sql:=dao.updatePhone
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,newPhone,phone)
	}
	da.Eq("phone",phone)
	da.Update("phone",newPhone)
	return  da.Execute()

}


/*
	根据手机号修改邮箱
*/
func (dao *BaseUserDaoImpl)UpdateEmailByPhone(phone string,email string)(int,error){
	sql:=dao.updateEmailByPhone
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,email,phone)
	}
	da.Eq("phone",phone)
	da.Update("email",email)
	return  da.Execute()
}

/*
	根据邮箱修改邮箱
*/
func (dao *BaseUserDaoImpl)UpdateEmail(email string,newEmail string)(int,error){
	sql:=dao.updateEmail
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,newEmail,email)
	}
	da.Eq("email",email)
	da.Update("email",newEmail)
	return  da.Execute()
}

/*
	根据旧密码修改密码
*/
func (dao *BaseUserDaoImpl)UpdatePwdByOldPwd(pwd string,newPwd string)(int,error){
	sql:=dao.updatePwd
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,newPwd,pwd)
	}
	da.Eq("pwd",pwd)
	da.Update("pwd",newPwd)
	return  da.Execute()
}

/*
	根据手机号修改密码
*/
func (dao *BaseUserDaoImpl)UpdatePwdByPhone(phone string,pwd string)(int,error){
	sql:=dao.updatePwdByPhone
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,pwd,phone)
	}
	da.Eq("phone",phone)
	da.Update("pwd",pwd)
	return  da.Execute()
}

/*
	根据邮箱修改密码
*/
func (dao *BaseUserDaoImpl)UpdatePwdByEmail(email string,pwd string)(int,error){
	sql:=dao.updatePwdByEmail
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,pwd,email)
	}
	da.Eq("email",email)
	da.Update("pwd",pwd)
	return  da.Execute()
}
/*
	修改登录次数
*/
func(dao *BaseUserDaoImpl) UpdateLoginFailCount(id int64)(int,error){
	sql:=dao.updateLoginFailCount
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,id)
	}
	da.Eq("id",id)
	da.UpdateIncr("login_fail_count")
	return  da.Execute()
}
func(dao *BaseUserDaoImpl) GetLoginFailCount(id int64)(int,error){
	sql:=dao.getLoginFailCount
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		da.ExecuteSql(sql,id)
		return  da.Count()
	}
	da.Eq("id",id)

	return  da.Count()
}

func(dao *BaseUserDaoImpl) ResetLoginFailCount(id int64)(int,error){
	sql:=dao.resetLoginFailCount
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,id)
	}
	da.Eq("id",id)
	da.Update("login_fail_count",0)
	return  da.Execute()
}

/*
	修改身份认证基本信息
*/
func(dao *UserDaoImpl)UpdateAuthBasic(input *dtos.UpdateUserAuthBasicInput)(int,error){
	sql:=dao.updateBasic
	da:=dao.Data()
	dao.init(da)
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return  da.ExecuteSqlToInt(sql,input.CardId,input.CardPhoto1,input.CardPhoto2,input.HandCardPhoto1,input.CardPhoto2,input.OperatorIp,input.OperatorDate,input.Id)
	}
	da.Eq("id",input.Id)
	da.Update("card_id",input.CardId)
	da.Update("card_photo1",input.CardPhoto1)
	da.Update("card_photo2",input.CardPhoto2)
	da.Update("hand_card_photo1",input.HandCardPhoto1)
	da.Update("hand_card_photo2",input.CardPhoto2)
	da.Update("update_ip",input.OperatorIp)
	da.Update("updates",input.OperatorDate)
	return  da.Execute()
}