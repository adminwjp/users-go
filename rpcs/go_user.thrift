//thrift-0.16.0.exe -gen java go_user.thrift
//thrift-0.16.0.exe -gen go go_user.thrift
//thrift-0.16.0.exe -gen netstd go_user.thrift

//namespace netstd com.user_system.ThriftRpcServices
namespace java com.mall.user.rpc.thrift.gouser
//namespace cpp com.user_system.thriftRpcs.services
//https://blog.csdn.net/weixin_30443075/article/details/99351633
//server is not null
namespace go users
/**
根据手机号、邮箱、用户名注册 或 登录实体
*/
struct UserInputThrift {
  1:string account //账号
  2:string pwd //密码
  3:i32 flag //账号 1 phone 2 email 3 user_name
}

/**
根据手机号注册 或 登录实体
*/
struct UserPhoneInputThrift {
  1:string phone //手机号
  2:string pwd //密码
}

/**
根据邮箱注册 或 登录实体
*/
struct UserEmailInputThrift {
  1:string email //邮箱
  2:string pwd //密码
}

/**
根据用户名注册 或 登录实体
*/
struct UserUserNameInputThrift {
  1:string user_name //用户名
  2:string pwd //密码
}

/**
根据手机号修改手机号
*/
struct UpdatePhoneInputThrift {
    1:string phone //手机号
    2:string new_phone //新手机号
}

/**
根据邮箱修改邮箱
*/
struct UpdateEmailInputThrift {
    1:string  email //邮箱
    2:string new_email //新邮箱
}

/**
根据手机号修改邮箱
*/
struct UpdateEmailByPhoneInputThrift {
    1:string  email //邮箱
    2:string phone //手机号
}

/**
根据邮箱修改密码
*/
struct UpdatePwdByEmailInputThrift {
    1:string email //邮箱
    2:string pwd //密码
}

/**
根据手机号修改密码
*/
struct UpdatePwdByPhoneInputThrift {
    1:string phone //手机号
    2:string pwd //密码
}

/**
根据旧密码修改密码
*/
struct UpdatePwdInputThrift {
    1:string pwd //密码
    2:string new_pwd //新密码
}

struct UpdateAuthBasicInputThrift {
    1:i64 user_id
    2:string card_id //身份证号
    3:string card_photo1 //身份证正面
    4:string card_photo2 //身份证反面
    5:string hand_card_photo1 //手持身份证正面
    6:string hand_card_photo2 //手持身份证反面
}

/**
根据手机号、邮箱、用户名检测账号是否存在
*/
struct AccountInputThrift {
  1:string account //账号
  2:i32 flag //账号 1 phone 2 email 3 user_name
}

/**
根据手机号检测账号是否存在
*/
struct PhoneInputThrift {
  1:string phone //手机号
}

/**
根据邮箱检测账号是否存在
*/
struct EmailInputThrift {
  1:string email //邮箱
}

/**
根据用户名检测账号是否存在
*/
struct UserNameInputThrift {
  1:string user_name //用户名
}

struct IdThrift {
  1:i64 id
}

struct ExistsOuputThrift {
  1:bool exists
}

struct UserOuputThrift {
    1:i64  id
    2:string phone
    3:string email
    4:string     user_name
    5:string pwd
    6:i64     reg_ip
    7:i64 login_ip
    8:i64     reg_date
    9:i64  login_date
    10:string card_id //身份证号
    11:string card_photo1 //身份证正面
    12:string card_photo2 //身份证反面
    13:string hand_card_photo1  //手持身份证正面
    14:string hand_card_photo2 //手持身份证反面
}
struct UserBasicOuputThrift {
    1:i64  id
    2:i64 user_id
    3:string card_id //身份证号
    4:string card_photo1 //身份证正面
    5:string card_photo2 //身份证反面
    6:string hand_card_photo1  //手持身份证正面
    7:string hand_card_photo2 //手持身份证反面
}
struct AdminOuputThrift {
    1:i64  id
    2:string phone
    3:string email
    4:string     user_name
    5:string pwd
    6:i64     reg_ip
    7:i64 login_ip
    8:i64     reg_date
    9:i64  login_date
}

struct ResultOuputThrift {
  1:i32 result
}

/*
用户 thrift 服务
*/
service UserThriftService {
	/*
	根据手机号、邮箱、用户名登录
	*/
    UserOuputThrift login (1:UserInputThrift user),

    /*
    	根据手机号登录
    */
    UserOuputThrift login_by_phone (1:UserPhoneInputThrift user),

 	/*
 	根据邮箱登录
 	*/
    UserOuputThrift login_by_email (1:UserEmailInputThrift user),

    /*
    根据用户名登录
    */
    UserOuputThrift login_by_user_name (1:UserUserNameInputThrift user),

    /*
    	根据手机号、邮箱、用户名注册 register 关键字 不能使用
    */
    IdThrift register1 (1:UserInputThrift user),

     /*
        	根据手机号注册
    */
    IdThrift register_by_phone (1:UserPhoneInputThrift user),

    /*
    根据邮箱注册
    */
    IdThrift register_by_email (1:UserEmailInputThrift user),

    /*
    根据用户名注册
    */
    IdThrift register_by_user_name (1:UserUserNameInputThrift user),

    /*
    根据id 获取用户信息
    */
    UserOuputThrift get (1:IdThrift user) ,

   UpdateAuthBasicInputThrift get_auth_basic (1:IdThrift user) ,

    /*
    		根据手机号修改手机号
    	*/
    ResultOuputThrift update_phone (1:UpdatePhoneInputThrift update) ,

    /*
    		根据邮箱修改邮箱
    */
    ResultOuputThrift update_email (1: UpdateEmailInputThrift update),

    /*
 		根据手机号修改邮箱
 	*/
    ResultOuputThrift update_email_by_phone (1:UpdateEmailByPhoneInputThrift update),

    /*
    		根据手机号修改密码
    	*/
    ResultOuputThrift update_pwd_by_phone (1:UpdatePwdByPhoneInputThrift update) ,

    /*
		根据邮箱修改密码
	*/
    ResultOuputThrift update_pwd_by_email (1:UpdatePwdByEmailInputThrift update) ,

    ResultOuputThrift update_auth_basic (1:UpdateAuthBasicInputThrift update),

/**
	根据手机号、邮箱、用户名检测账号是否存在
	*/
    ExistsOuputThrift exists (1:AccountInputThrift user),

    /**
    	根据手机号检测账号是否存在
    	*/
    ExistsOuputThrift exists_phone (1:PhoneInputThrift user),

    /**
    根据邮箱检测账号是否存在
    */
    ExistsOuputThrift exists_email (1:EmailInputThrift user),

   	/**
   	根据用户名检测账号是否存在
   	*/
   ExistsOuputThrift exists_user_name (1:UserNameInputThrift user)
}

/*
管理员 thrift 服务
*/
service AdminThriftService {
	/*
	根据手机号、邮箱、用户名登录
	*/
    AdminOuputThrift login (1:UserInputThrift user),

    /*
    	根据手机号登录
    */
    AdminOuputThrift login_by_phone (1:UserPhoneInputThrift user),

 	/*
 	根据邮箱登录
 	*/
    AdminOuputThrift login_by_email (1:UserEmailInputThrift user),

    /*
    根据用户名登录
    */
    AdminOuputThrift login_by_user_name (1:UserUserNameInputThrift user),

    /*
    	根据手机号、邮箱、用户名注册 register 关键字 不能使用
    */
    IdThrift register1 (1:UserInputThrift user),

     /*
        	根据手机号注册
    */
    IdThrift register_by_phone (1:UserPhoneInputThrift user),

    /*
    根据邮箱注册
    */
    IdThrift register_by_email (1:UserEmailInputThrift user),

    /*
    根据用户名注册
    */
    IdThrift register_by_user_name (1:UserUserNameInputThrift user),

    /*
    根据id 获取用户信息
    */
    AdminOuputThrift get (1:IdThrift user) ,

    /*
    		根据手机号修改手机号
    	*/
    ResultOuputThrift update_phone (1:UpdatePhoneInputThrift update) ,

    /*
    		根据邮箱修改邮箱
    */
    ResultOuputThrift update_email (1:UpdateEmailInputThrift update),

    /*
 		根据手机号修改邮箱
 	*/
    ResultOuputThrift update_email_by_phone (1:UpdateEmailByPhoneInputThrift update),

    /*
    		根据手机号修改密码
    	*/
    ResultOuputThrift update_pwd_by_phone (1:UpdatePwdByPhoneInputThrift update) ,

    /*
		根据邮箱修改密码
	*/
    ResultOuputThrift update_pwd_by_email (1:UpdatePwdByEmailInputThrift update),

/*
		根据旧密码修改密码
	*/
    ResultOuputThrift update_pwd (1:UpdatePwdInputThrift update),

/**
	根据手机号、邮箱、用户名检测账号是否存在
	*/
    ExistsOuputThrift exists (1:AccountInputThrift user),

    /**
    	根据手机号检测账号是否存在
    	*/
    ExistsOuputThrift exists_phone (1:PhoneInputThrift user),

    /**
    根据邮箱检测账号是否存在
    */
    ExistsOuputThrift exists_email (1:EmailInputThrift user),

   	/**
   	根据用户名检测账号是否存在
   	*/
   ExistsOuputThrift exists_user_name (1:UserNameInputThrift user)
}