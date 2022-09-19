

1.用户接口



1.2 用户注册

1.2.1 根据手机号、邮箱、用户名注册

##### rpc method
   register


|  name   | type   | length | comment                    |
| :-----: | ------ | ------ | -------------------------- |
| account | string |        | 手机号、邮箱、用户名       |
|   pwd   | string |        | 密码                       |
|  flag   | int    |        | 1:手机号、2:邮箱、3:用户名 |




1.2.2 根据手机号注册

##### rpc method
   register_phone

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| phone | string | 11     | 手机号  |
|  pwd  | string |        | 密码    |



1.2.3 根据邮箱注册

##### rpc method
   register_email


| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |
|  pwd  | string |        | 密码    |



1. 2.4根据用户名注册

##### rpc method
   register_user_name


|   name    | type   | length | comment |
| :-------: | ------ | ------ | ------- |
| user_name | string | 20     | 用户名  |
|    pwd    | string |        | 密码    |

