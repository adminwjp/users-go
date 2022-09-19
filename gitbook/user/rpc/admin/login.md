1.1 用户登录

1.1.1 根据手机号、邮箱、用户名登录

##### rpc method
   login

|   name    | type   | length | comment                    |
| :-------: | ------ | ------ | -------------------------- |
|  account  | string |        | 手机号、邮箱、用户名       |
|    pwd    | string |        | 密码                       |
|   flag    | int    |        | 1:手机号、2:邮箱、3:用户名 |



1.1.2 根据手机号登录

##### rpc method
   login_phone

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| phone | string | 11     | 手机号  |
|  pwd  | string |        | 密码    |



1.1.3 根据邮箱登录

##### rpc method
   login_email

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |
|  pwd  | string |        | 密码    |



1. 1.4根据用户名登录

##### rpc method
   login_user_name

|   name    | type   | length | comment |
| :-------: | ------ | ------ | ------- |
| user_name | string | 20     | 用户名  |
|    pwd    | string |        | 密码    |

