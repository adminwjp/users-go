1.6修改密码

1.6.1根据邮箱修改密码



##### post

​	http://www.user.com/admin/update/pwd/email



​	http://www.user.com/admin/update_pwd/email



​	http://www.user.com/admin/update_pwd_by_email



| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |
|  pwd  | string |        | 密码    |

###### post json

```json
{"email":"123@qq.com","pwd":"1234"}
```

##### post xml

```xml
<user>
    <email>123@qq.com</email>
    <pwd>1234</pwd>
</user>
```

1.6.2根据手机号修改密码


##### post

​	http://www.user.com/admin/update/pwd/phone



http://www.user.com/admin/update_pwd/phone



http://www.user.com/admin/update_pwd_by_phone



| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
|  pwd  | string |        | 密码    |
| phone | string | 11     | 手机号  |

###### post json

```json
{"phone":"13147444320","pwd":"1234"}
```

##### post xml

```xml
<user>
    <phone>13147444320</phone>
    <pwd>1234</pwd>
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

1.6.3根据密码修改密码



##### post

​	http://www.user.com/admin/update/pwd



    http://www.user.com/admin/update_pwd




| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
|  pwd  | string |        | 密码    |
| new_pwd | string |     | 新密码  |

###### post json

```json
{"pwd":"123","new_pwd":"1234"}
```

##### post xml

```xml
<user>
    <pwd>123</pwd>
    <new_pwd>1234</new_pwd>
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