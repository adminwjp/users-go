1.5修改邮箱

1.5.1根据邮箱修改邮箱

##### get

​	http://www.user.com/admin/update/email/123@qq.com/1233@qq.com

http://www.user.com/admin/update/email?email=123@qq.com&new_email=1233@qq.com



​	http://www.user.com/admin/update_email/123@qq.com/1233@qq.com

http://www.user.com/admin/update_email?email=123@qq.com&new_email=1233@qq.com

##### post

​	http://www.user.com/admin/update/email



http://www.user.com/admin/update_email

|   name    | type   | length | comment |
| :-------: | ------ | ------ | ------- |
|   email   | string | 20     | 邮箱    |
| new_email | string | 20     | 新邮箱  |

###### post json

```json
{"email":"123@qq.com","new_email":"1233@qq.com"}
```

##### post xml

```xml
<user>
    <email>123@qq.com</email>
    <new_email>1233@qq.com</new_email>
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

1.5.2根据手机号修改邮箱

##### get

​    http://www.user.com/admin/update/email/phone/13147444320/1233@qq.com

​	http://www.user.com/admin/update/email/phone?email=13147444320&new_email=1233@qq.com



​	http://www.user.com/admin/update_email/phone/13147444320/1233@qq.com

​	http://www.user.com/admin/update_email/phone?email=13147444320&new_email=1233@qq.com



​	http://www.user.com/admin/update_emai_by_phone/13147444320/1233@qq.com

​	http://www.user.com/admin/update_emai_by_phone?email=13147444320&new_email=1233@qq.com

##### post

​	http://www.user.com/admin/update_email/phone

​	http://www.user.com/admins/update/email/phone

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |
| phone | string | 11     | 手机号  |

###### post json

```json
{"phone":"13147444320","email":"1233@qq.com"}
```

##### post xml

```xml
<user>
    <phone>13147444320</phone>
    <email>1233@qq.com</email>
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