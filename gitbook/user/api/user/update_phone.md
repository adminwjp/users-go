1.4修改手机号

##### get

​	http://www.user.com/user/update/phone/13147444320/13147444321

​	http://www.user.com/user/update/phone?phone=13147444320&new_phone=13147444321



##### post

​	http://www.user.com/user/update/phone

|   name    | type   | length | comment  |
| :-------: | ------ | ------ | -------- |
|   phone   | string | 11     | 手机号   |
| new_phone | string | 11     | 新手机号 |

###### post json

```json
{"phone":"13147444320","new_phone":"13147444321"}
```

##### post xml

```xml
<user>
    <phone>13147444320</phone>
    <new_phone>13147444321</new_phone>
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