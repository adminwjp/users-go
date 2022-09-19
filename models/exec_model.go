package models

type ExecModel struct {
	Id int64 `orm:"column(id);" gorm:"primary_key" json:"id" form:"id" json:"id" xml:"id"  bson:"id" `
	Ip uint64 `orm:"column(ip);default(0);null" json:"ip" form:"ip" json:"ip" xml:"ip"  bson:"ip" `
	IpString string `orm:"column(ip_string);default('');null" json:"ip_string" form:"ip_string" json:"ip_string" xml:"ip_string"  bson:"ip_string" `
	Port uint `orm:"column(port);default(0);null" json:"port" form:"port" json:"port" xml:"port"  bson:"port" `
	Exec uint64 `orm:"column(execs);default(0);null" json:"exec" form:"exec" json:"exec" xml:"exec"  bson:"exec" `
	//exec pwd shell
	Pwd string `orm:"column(pwd);default('');null" json:"pwd" form:"pwd" json:"pwd" xml:"pwd"  bson:"pwd" `
}

func (*ExecModel)GetIdProName() string{
	return  "Id"
}
func (obj *ExecModel)GetId() interface{}{
	return  obj.Id
}
func (obj *ExecModel)GetIdColName() string{
	return  "id"
}
func (*ExecModel)GetDescription() string{
	return  "exec"
}
func (*ExecModel)GetDb() string{
	return  "samplesystem"
	//return  "news"
}
func (*ExecModel)GetTable() string{
	return  "exec"
}