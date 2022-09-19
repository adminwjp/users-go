1.3 检测用户是否存在

1.3.1 检测手机号、邮箱、用户名是否存在

##### rpc method

​	exists


|  name   | type   | length | comment                    |
| :-----: | ------ | ------ | -------------------------- |
| account | string |        | 手机号、邮箱、用户名       |
|  flag   | int    |        | 1:手机号、2:邮箱、3:用户名 |

1.3.2 检测手机号是否存在

##### rpc method
   exists_phone

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| phone | string | 11     | 手机号  |

1.3.3 检测邮箱是否存在

##### rpc method
   exists_email


| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |

1. 3.4检测用户名是否存在

##### rpc method
   exists_user_name

|   name    | type   | length | comment |
| :-------: | ------ | ------ | ------- |
| user_name | string | 20     | 用户名  |