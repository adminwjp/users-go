package daos

import (
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/users-go/dtos"
	"log"
	"time"
)

func(dao *BaseUserDaoImpl)GetTranction()TranDao{
	return dao.TranManager
}

type BaseUserDaoImpl struct {
	M interface{}
	TranManager TranDao
	//db sql
	Data func()Dao
	registerByPhone string
	registerByEmail string
	registerByUserName string

	user string
	isAdmin bool


	existsByPhone string
	existsByEmail string
	existsByUserName string

	updatePhone string

	updateEmail string
	updateEmailByPhone string

	updatePwd string
	updatePwdByEmail string
	updatePwdByPhone string

	updateBasic string

	selectAll string

	loginByPhone string
	loginByEmail string
	loginByUserName string

	updateLoginFailCount string
	resetLoginFailCount string

	getLoginFailCount string

	view string
	table string
	doc string
	collection string
}


func (dao *BaseUserDaoImpl)SetAdmin(admin bool){
	dao.isAdmin=admin
}
func (dao *BaseUserDaoImpl)SetUserSql()  {
	dao.user="admin"
	view:="v_admin"
	table:="t_admin"
	if !dao.isAdmin{
		dao.user="user"
		view="v_user"
		table="t_user"
	}
	dao.view=view
	dao.table=table
	dao.doc=dao.user
	dao.collection=dao.user

	dao.registerByPhone="insert into "+table+" (id,phone,pwd,reg_ip,reg_date) values(?,?,?,?,?)"
	dao.registerByEmail="insert into "+table+" (id,email,pwd,reg_ip,reg_date) values(?,?,?,?,?)"
	dao.registerByUserName="insert into "+table+" (id,user_name,pwd,reg_ip,reg_date) values(?,?,?,?,?)"

	//防止数据一直更新 加当前时间戳
	dao.existsByPhone="select count(*) total from "+view+" where phone=? and reg_date <"//?"
	dao.existsByEmail="select count(*) total from "+view+" where email=? and reg_date <"//?"
	dao.existsByUserName="select count(*) total from "+view+" where user_name= ? and reg_date <"//?"

	dao.updatePhone="update "+table+" set phone=? where phone=?"

	dao.updateEmailByPhone="update "+table+" set email=? where phone=?"
	dao.updateEmail="update "+table+" set email=? where email=?"

	dao.updatePwd="update "+table+" set pwd=? where pwd=?"
	dao.updatePwdByEmail="update "+table+" set pwd=? where email=?"
	dao.updatePwdByPhone="update "+table+" set pwd=? where phone=?"

	dao.updateLoginFailCount="update "+table+" set login_fail_count=login_fail_count+1 where id=?"

	dao.updateBasic="update "+table+" set card_id=?,card_photo1=?,card_photo2=?,hand_card_photo1=?,hand_card_photo2=?,update_ip=?,updates=? where id=?"

	dao.resetLoginFailCount="update "+table+" set login_fail_count=0 where id=?"
	dao.getLoginFailCount="select login_fail_count from  "+view+"  where id=?"


	dao.selectAll="select * from "+view+""

	//防止数据一直更新 加当前时间戳
	dao.loginByPhone="select * from "+view+" where phone=? and pwd=? and reg_date <"//?"
	dao.loginByEmail="select * from "+view+" where email=? and pwd=? and reg_date <"//?"
	dao.loginByUserName="select * from "+view+" where user_name=? and pwd=? and reg_date <"//?"
}

func(dao *BaseUserDaoImpl) Get(id int64,obj interface{})error{
	da:=dao.Data()
	dao.init(da)

	da.Eq("id",id)
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)
	da.Lt("reg_date",n.Unix())
	return da.One(obj)
}
func (dao *BaseUserDaoImpl)init(da Dao)  {
	da.View(dao.view)
	da.Doc(dao.doc)
	da.Db(MDb)
	log.Println("set collection => "+dao.collection)
	da.Collection(dao.collection)
}

func(dao *BaseUserDaoImpl) List(user *dtos.GetBaseUserInput,list interface{}) (int64,error){
	da:=dao.Data()
	dao.init(da)
	//db=db.Model(UserModel{})
	if user.Phone!=""{
		da.OrLike("phone",user.Phone)
	}
	if user.Email!=""{
		da.OrLike("email",user.Email)
	}
	if user.UserName!=""{
		da.OrLike("user_name",user.UserName)
	}
	if user.RegIp>0{
		da.OrEq("reg_ip",user.RegIp)
	}
	if user.LoginIp>0{
		da.OrEq("login_ip",user.LoginIp)
	}
	n:=time.Now()
	n=n.Add(-time.Millisecond*30)

	/*if user.LoginStartDate>0&& user.LoginEndDate>0{
		if user.LoginStartDate<n.Unix(){
			if user.LoginEndDate<n.Unix(){
				db=db.Or(" login_date between ? and ? ",user.LoginStartDate,user.LoginEndDate)
			}else{
				db=db.Or(" login_date between ? and ? ",user.LoginStartDate,n.Unix())
			}
		}
	}else if user.LoginStartDate>0{
		if user.LoginStartDate>=n.Unix(){

		}else{
			db=db.Or(" login_date between ? and ? ",user.LoginStartDate,n.Unix())
		}
	}else if user.LoginEndDate>0{
		if user.LoginEndDate>int64(n.Unix()){
			db=db.Or(" login_date < ? ",n.Unix())
		}else{
			db=db.Or(" login_date < ? ",user.LoginEndDate)
		}
	}*/

	if user.RegStartDate>0&& user.RegEndDate>0{
		n:=time.Now()
		n=n.Add(-time.Millisecond*30)
		if user.RegStartDate<n.Unix(){
			if user.RegEndDate<n.Unix(){
				da.OrBetween("reg_date",user.RegStartDate,user.RegEndDate)
			}else{
				da.OrBetween("reg_date",user.RegStartDate,n.Unix())
			}
		}

	}else if user.RegStartDate>0{
		n:=time.Now()
		n=n.Add(-time.Millisecond*30)
		//	db=db.Or(" reg_date > ?  ",user.RegStartDate)
		if user.RegStartDate>=n.Unix(){

		}else{
			da.OrBetween("reg_date",user.RegStartDate,n.Unix())
		}
	}else if user.RegEndDate>0{
		n:=time.Now()
		n=n.Add(-time.Millisecond*30)
		if user.RegEndDate>int64(n.Unix()){
			da.OrLt("reg_date",n.Unix())
		}else{
			da.OrLt("reg_date",user.RegEndDate)
		}
	}
	//da.View(dao.view)
	da.Model(dao.M)
	return da.ListByPage(user.Page,user.Size,list)
}

type BaseUserDao interface {
	Init()
	GetTranction() TranDao
	//SetTranction(tran TranDao)
	/*
		根据手机号、邮箱、用户名注册
	*/
	Register(user *dtos.UserInput)(int,error)


	/*
		根据手机号注册
	*/
	RegisterByPhone(user *dtos.UserPhoneInput)(int,error)

	/*
		根据邮箱注册
	*/
	RegisterByEmail(user *dtos.UserEmailInput)(int,error)

	/*
		根据用户名注册
	*/
	RegisterByUserName(user *dtos.UserUserNameInput)(int,error)

	/*
		根据手机号、邮箱、用户名检测账号是否存在
	*/
	Exists(account string,flag dto.AccounType)(int,error)


	/*
		根据手机号测账号是否存在
	*/
	ExistsByPhone(phone string)(int,error)
	/*
		根据邮箱测账号是否存在
	*/
	ExistsByEmail(email string)(int,error)

	/*
		根据用户名测账号是否存在
	*/
	ExistsByUserName(userName string)(int,error)

	/*
		根据手机号修改手机号
	*/
	UpdatePhone(phone string,newPhone string)(int,error)


	/*
		根据手机号修改邮箱
	*/
	UpdateEmailByPhone(phone string,email string)(int,error)

	/*
		根据邮箱修改邮箱
	*/
	UpdateEmail(email string,newEmail string)(int,error)

	/*
		根据手机号修改密码
	*/
	UpdatePwdByPhone(phone string,pwd string)(int,error)

	/*
		根据邮箱修改密码
	*/
	UpdatePwdByEmail(email string,pwd string)(int,error)
	/*
		修改登录次数
	*/
	 UpdateLoginFailCount(id int64)(int,error)
	 GetLoginFailCount(id int64)(int,error)

	 ResetLoginFailCount(id int64)(int,error)

}
