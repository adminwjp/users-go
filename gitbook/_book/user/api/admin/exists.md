1.3 检测用户是否存在

1.3.1 检测手机号、邮箱、用户名是否存在

##### get

​	http://www.user.com/admin/exists/123/3

​	http://www.user.com/admin/exists?account=123&flag=3

|  name   |  type   |  length  |   comment                    |
| :-----: | :-----: | :------: | :--------------------------: |
| account |  string |          |    手机号、邮箱、用户名       |
|  flag   |  int    |          |   1:手机号、2:邮箱、3:用户名 |

1.3.2 检测手机号是否存在

​	http://www.user.com/admin/phone/exists/13147444320

​	http://www.user.com/admin/phone/exists?phone=13147444320

| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| phone | string | 11     | 手机号  |

##### response string
```string
false
```
##### response xml
```xml
<response>false</response>
```

1.3.3 检测邮箱是否存在

##### get

​	http://www.user.com/admin/email/exists/123@qq.com

​	http://www.user.com/admin/email/exists?email=123@qq.com


| name  | type   | length | comment |
| :---: | ------ | ------ | ------- |
| email | string | 20     | 邮箱    |

##### response string
```string
false
```
##### response xml
```xml
<response>false</response>
```

1. 3.4检测用户名是否存在

   ##### get

​	http://www.user.com/admin/user_name/exists/123

   http://www.user.com/admin/user_name/exists?user_name=123

|   name    |   type   |   length |   comment |
| :-------: | :------: | :------: | :-------: |
| user_name |   string |   20     |   用户名  |

##### response string
```string
false
```
##### response xml
```xml
<response>false</response>
```