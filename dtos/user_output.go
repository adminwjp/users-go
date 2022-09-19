package dtos

type GetUserOutput struct {
	//用户id
	Id int64 `orm:"column(id)" gorm:"id;primary_key" json:"id" form:"id"  xml:"id" bson:"id"`

	//手机号
	Phone string `orm:"column(phone)" gorm:"phone" json:"phone" form:"phone"  xml:"phone" bson:"phone"`

	//账号
	//Account string `orm:"column(account)" gorm:"account" json:"account" form:"account"  xml:"account" `

	//用户名
	UserName string `orm:"column(user_name)" gorm:"user_name" json:"user_name" form:"user_name"  xml:"user_name" bson:"user_name" `

	//邮箱
	Email string `orm:"column(email)" gorm:"email" json:"email" form:"email"  xml:"email" bson:"email"`

	//密码
	Pwd string `orm:"column(pwd)" gorm:"pwd" json:"pwd" form:"pwd"  xml:"pwd" bson:"pwd"`

	//注册Ip
	RegIp int64 `orm:"column(reg_ip)" gorm:"reg_ip" json:"reg_ip" form:"reg_ip"  xml:"reg_ip" bson:"reg_ip"`

	//登录Ip
	// LoginIp int64 `orm:"column(login_ip)" gorm:"login_ip" json:"login_ip" form:"login_ip"  xml:"login_ip" bson:"login_ip" `

	//注册时间
	RegDate int64 `orm:"column(reg_date)" gorm:"reg_date;default:0" json:"-" form:"-"  xml:"-" bson:"reg_date"`

	//登录时间
	// LoginDate int64 `orm:"column(login_date)" gorm:"login_date;default:0" json:"login_date" form:"login_date"  xml:"login_date" bson:"login_date" `

	//修改时间
	// UpdateDate int64 `orm:"column(update_date)" gorm:"update_date;default:0" json:"update_date" form:"update_date"  xml:"update_date" bson:"update_date" `

	Flag string `orm:"column(flag)" gorm:"flag" json:"-" form:"-"  xml:"-" bson:"flag"`

	LoginFailCount int `orm:"column(login_fail_count)" gorm:"login_fail_count" json:"-" form:"-"  xml:"-" bson:"login_fail_count"`

	//昵称
	NickName string `orm:"column(nick_name)" gorm:"nick_name" json:"nick_name" form:"nick_name"  xml:"nick_name" bson:"nick_name"`

	//角色Id
	RoleId int64 `orm:"column(role_id)" gorm:"role_id" json:"role_id" form:"role_id"  xml:"role_id" bson:"role_id"`
}
