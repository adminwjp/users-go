package daos

import (
	data "github.com/adminwjp/infrastructure-go/datas"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/datas"
	"github.com/adminwjp/users-go/dtos"
)

func (dao *BaseUserDaoImpl)Register(input *dtos.UserInput,method func(*dtos.UserInput)interface{})(int,error){
	//reg reg_update_phone_update_email -> rpe
	da:=dao.Data()
	dao.init(da)
	sql:=dao.registerByPhone
	if input.Flag == dto.AccounTypeByEamil {
		sql = dao.registerByEmail
	} else if input.Flag == dto.AccounTypeByUsername {
		sql =dao.registerByUserName
	}
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return da.ExecuteSqlToInt(sql,input.Id,input.Account,input.Pwd,input.OperatorIp,input.OperatorDate)
	}
	return da.Add(method(input))
}
func (dao *BaseUserDaoImpl)RegisterByPhone(input *dtos.UserPhoneInput,method func(*dtos.UserPhoneInput)interface{})(int,error){
	//reg reg_update_phone_update_email -> rpe
	da:=dao.Data()
	dao.init(da)
	sql:=dao.registerByPhone
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return da.ExecuteSqlToInt(sql,input.Id,input.Phone,input.Pwd,input.OperatorIp,input.OperatorDate)
	}
	return da.Add(method(input))
}

func (dao *BaseUserDaoImpl)RegisterByEmail(input *dtos.UserEmailInput,method func(*dtos.UserEmailInput)interface{})(int,error){
	//reg reg_update_phone_update_email -> rpe
	da:=dao.Data()
	dao.init(da)
	sql:= dao.registerByEmail

	if datas.GlobalConfig.DataFlag== data.DataDb{
		return da.ExecuteSqlToInt(sql,input.Id,input.Email,input.Pwd,input.OperatorIp,input.OperatorDate)
	}
	return da.Add(method(input))
}

func (dao *BaseUserDaoImpl)RegisterByUserName(input *dtos.UserUserNameInput,method func(*dtos.UserUserNameInput)interface{})(int,error){
	//reg reg_update_phone_update_email -> rpe
	da:=dao.Data()
	dao.init(da)
	sql:=dao.registerByUserName
	if datas.GlobalConfig.DataFlag== data.DataDb{
		return da.ExecuteSqlToInt(sql,input.Id,input.UserName,input.Pwd,input.OperatorIp,input.OperatorDate)
	}
	return da.Add(method(input))
}