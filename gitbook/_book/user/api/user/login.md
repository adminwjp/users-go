1.1 用户登录

1.1.1 根据手机号、邮箱、用户名登录

###### post

​	http://www.user.com/user/login

|   name    | type   | length | comment                    |
| :-------: | ------ | ------ | -------------------------- |
|  account  | string |        | 手机号、邮箱、用户名       |
|    pwd    | string |        | 密码                       |
|   flag    | int    |        | 1:手机号、2:邮箱、3:用户名 |

###### post json
```json
{"account":"123","pwd":"123","flag":3}
```

##### post xml

```xml
<user>
	<account>123</account>
	<pwd>123</pwd>
	<flag>3</flag>
</user>
```

##### response json
```json
{
    
}
```
##### response xml
```xml
<response>
</response>
```

1.1.2 根据手机号登录

###### post

​	http://www.user.com/user/phone/login

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| phone | string | 11     | 手机号  |
|  pwd  | string |        | 密码    |

###### post json
```json
{"phone":"13147444320","pwd":"123"}
```

##### post xml

```xml
<user>
    <phone>13147444320</phone>
    <pwd>123</pwd>
</user>
```

##### response json
```json
{
    
}
```
##### response xml
```xml
<response>
</response>
```

1.1.3 根据邮箱登录

###### post

​	http://www.user.com/user/email/login

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |
|  pwd  | string |        | 密码    |

###### post json

```json
{"email":"123@qq.com","pwd":"123"}
```

##### post xml

```xml
<user>
    <email>123@qq.com</email>
    <pwd>123</pwd>
</user>
```
##### response json
```json
{
    
}
```
##### response xml
```xml
<response>
</response>
```

1. 1.4根据用户名登录

###### post

​	http://www.user.com/user/user_name/login

|   name    | type   | length | comment |
| :-------: | ------ | ------ | ------- |
| user_name | string | 20     | 用户名  |
|    pwd    | string |        | 密码    |

###### post json

```json
{"email":"123","pwd":"123"}
```

##### post xml

```xml
<user>
    <user_name>123</user_name>
    <pwd>123</pwd>
</user>
```

##### response json
```json
{
    
}
```
##### response xml
```xml
<response>
</response>
```